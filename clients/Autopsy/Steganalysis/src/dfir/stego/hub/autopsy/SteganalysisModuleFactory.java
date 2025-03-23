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

import org.openide.util.NbBundle;
import org.openide.util.lookup.ServiceProvider;
import org.sleuthkit.autopsy.coreutils.Version;
import org.sleuthkit.autopsy.ingest.FileIngestModule;
import org.sleuthkit.autopsy.ingest.IngestModuleFactory;
import org.sleuthkit.autopsy.ingest.IngestModuleFactoryAdapter;
import org.sleuthkit.autopsy.ingest.IngestModuleIngestJobSettings;
import org.sleuthkit.autopsy.ingest.IngestModuleIngestJobSettingsPanel;

/**
 * Factory for the Picture Analysis ingest module.
 */
@ServiceProvider(service = IngestModuleFactory.class)
public class SteganalysisModuleFactory extends IngestModuleFactoryAdapter {

    public static String getModuleName() {
        return NbBundle.getMessage(SteganalysisModuleFactory.class, "moduleDisplayName.text");
    }
    
    @Override
    public String getModuleDisplayName() {
        return getModuleName();
    }

    @Override
    public String getModuleDescription() {
        return NbBundle.getMessage(SteganalysisModuleFactory.class, "moduleDescription.text");
    }

    @Override
    public String getModuleVersionNumber() {
        return Version.getVersion();
    }
    
    @Override
    public boolean hasIngestJobSettingsPanel() {
        return true;
    }
    
    @Override
    public IngestModuleIngestJobSettingsPanel getIngestJobSettingsPanel(IngestModuleIngestJobSettings settings) {
        return new SteganalysisIngestJobSettingsPanel((SteganalysisIngestJobSettings) settings);
    }
    
    @Override
    public IngestModuleIngestJobSettings getDefaultIngestJobSettings() {
        return new SteganalysisIngestJobSettings();
    }

    @Override
    public FileIngestModule createFileIngestModule(IngestModuleIngestJobSettings ingestOptions) {
        return new SteganalysisModule((SteganalysisIngestJobSettings)ingestOptions);
    }
    
    @Override
    public boolean isFileIngestModuleFactory() {
        return true;
    }

}
