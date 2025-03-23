/*
 * Autopsy Forensic Browser
 *
 * Copyright 2020 Basis Technology Corp.
 * Contact: carrier <at> sleuthkit <dot> org
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package dfir.stego.hub.autopsy;

import com.google.protobuf.ByteString;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import java.util.Map;
import java.util.logging.Level;
import org.sleuthkit.autopsy.casemodule.Case;
import org.sleuthkit.autopsy.casemodule.NoCurrentCaseException;
import org.sleuthkit.autopsy.coreutils.Logger;
import org.sleuthkit.autopsy.coreutils.MessageNotifyUtil;
import org.sleuthkit.autopsy.ingest.FileIngestModuleAdapter;
import org.sleuthkit.autopsy.ingest.IngestJobContext;
import org.sleuthkit.autopsy.ingest.IngestModule;
import org.sleuthkit.autopsy.ingest.IngestModuleReferenceCounter;

import org.sleuthkit.datamodel.AbstractFile;
import org.sleuthkit.datamodel.Blackboard;
import org.sleuthkit.datamodel.BlackboardArtifact;
import org.sleuthkit.datamodel.TskCoreException;
import org.sleuthkit.datamodel.TskDataException;
import stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest;
import stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse;
import stego.hub.grpc.wrapper.StegAnalysis.TaskResult;
import stego.hub.grpc.wrapper.StegAnalysisServiceGrpc;
import stego.hub.grpc.wrapper.StegAnalysisServiceGrpc.StegAnalysisServiceBlockingStub;
import stego.hub.grpc.wrapper.StegServiceOuterClass;
import stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse;

/**
 * Ingest module to do arbitrary work on image files. Some examples include
 * extracting EXIF, converting between formats, and creating thumbnails. This
 * module acts as a container for multiple PictureProcessors, which are the
 * classes that do the work mentioned in the examples above.
 */
public class SteganalysisModule extends FileIngestModuleAdapter {

    private static final Logger logger = Logger.getLogger(SteganalysisModule.class.getName());
    private static final IngestModuleReferenceCounter refCounter = new IngestModuleReferenceCounter();
    private final SteganalysisIngestJobSettings settings;
    private IngestJobContext context;
    private ManagedChannel channel;
    private StegAnalysisServiceBlockingStub stub;
    private long jobId;

    public SteganalysisModule(SteganalysisIngestJobSettings settings) {
        this.settings = settings;
    }

    @Override
    public ProcessResult process(AbstractFile file) {

        try {
            if (!settings.getExtensions().contains(file.getNameExtension())) {
                return ProcessResult.OK;
            }

            long fileSize = file.getSize();

            if (fileSize > settings.getMaxFileSize() && settings.getMaxFileSize() != -1) {
                return ProcessResult.OK;
            }

            byte[] fileContent = new byte[(int) fileSize];
            file.read(fileContent, 0, fileContent.length);

            StegAnalysisRequest request = StegAnalysisRequest
                    .newBuilder()
                    .setFile(ByteString.copyFrom(fileContent))
                    .putAllParams(settings.getAdditionalParams())
                    .setFileName(file.getName())
                    .build();

            StegAnalysisResponse response = stub.execute(request);

            addStegAnalysisResponseToBlackboard(file, response);

        } catch (TskCoreException ex) {
            logger.log(Level.SEVERE, String.format("Unable to create artifact for file (obj_id=%d)", file.getId()), ex);
            MessageNotifyUtil.Notify.error(SteganalysisModuleFactory.getModuleName(), String.format("Unable to create TSK_INTERESTING_ITEM artifact for file (obj_id=%d)", file.getId()));
            return ProcessResult.ERROR;
        } catch (Exception ex) {
            logger.log(Level.WARNING, String.format("error while prcessing file '%s'.", file.getName()), ex); // NON-NLS
            MessageNotifyUtil.Notify.error(SteganalysisModuleFactory.getModuleName(), String.format("error while prcessing file '%s': '%s'", file.getName(), ex.getMessage()));
            return ProcessResult.ERROR;
        }

        return ProcessResult.OK;
    }

    public void addStegAnalysisResponseToBlackboard(AbstractFile file, StegAnalysisResponse response) throws TskDataException {
        try {

            BlackboardArtifact artifact = CustomArtifactHelper.addCustomArtifactToFile(file);

            if (!response.getError().equals("")) {
                CustomAttributeHelper.addCustomAttributeToArtifact(artifact, "ERROR", "Error", response.getError());
            }

            for (TaskResult taskResult : response.getTaskResultsList()) {
                if (!taskResult.getError().equals("")) {
                    CustomAttributeHelper.addCustomAttributeToArtifact(artifact, String.format("%s_ERROR", taskResult.getTaskId()), String.format("%s Error", taskResult.getTaskId()), taskResult.getError());
                }

                StegServiceResponse serviceResponse = taskResult.getServiceResponse();
                if (serviceResponse != null) {
                    if (!serviceResponse.getError().equals("")) {
                        CustomAttributeHelper.addCustomAttributeToArtifact(artifact, String.format("%s_ERROR", taskResult.getTaskId()), String.format("%s Error", taskResult.getTaskId()), serviceResponse.getError());
                    }
                    for (Map.Entry<String, StegServiceOuterClass.ResponseValue> value : serviceResponse.getValuesMap().entrySet()) {
                        CustomAttributeHelper.addCustomAttributeToArtifact(artifact, String.format("TASK_RETURN_VALUE_%s_%s", taskResult.getTaskId(), value.getKey()), String.format("%s->%s", taskResult.getTaskId(), value.getKey()), value.getValue().toString());
                    }
                }
            }

            Case currentCase = Case.getCurrentCaseThrows();
            Blackboard tskBlackboard = currentCase.getSleuthkitCase().getBlackboard();

            tskBlackboard.postArtifact(artifact, SteganalysisModuleFactory.getModuleName(), jobId);

        } catch (TskCoreException | Blackboard.BlackboardException ex) {
            logger.log(Level.SEVERE, "error adding artifact to blackboard.", ex);
        } catch (NoCurrentCaseException ex) {
            logger.log(Level.SEVERE, "exception while getting open case.", ex);
        }
    }

    @Override
    public void startUp(IngestJobContext context) throws IngestModuleException {
        try {
            this.context = context;
            jobId = context.getJobId();
            refCounter.incrementAndGet(jobId);

            if (channel != null) {
                return;
            }

            String grpcServerAddress = settings.getEndpoint().split(":")[0];
            int grpcPort = Integer.parseInt(settings.getEndpoint().split(":")[1]);

            channel = ManagedChannelBuilder
                    .forAddress(grpcServerAddress, grpcPort)
                    .usePlaintext()
                    .build();
            stub = StegAnalysisServiceGrpc.newBlockingStub(channel);

            logger.log(Level.SEVERE, "connected to dfir stego-hub");
        } catch (NumberFormatException e) {
            throw new IngestModule.IngestModuleException("could not parse endpoint:", e);
        } catch (Exception e) {
            throw new IngestModule.IngestModuleException("could not connect to endpoint:", e);
        }
    }

    @Override
    public void shutDown() {
        if (channel != null) {
            channel.shutdown();
        }
    }
}
