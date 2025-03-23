/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package dfir.stego.hub.autopsy;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import org.sleuthkit.autopsy.ingest.IngestModuleIngestJobSettings;

public class SteganalysisIngestJobSettings implements IngestModuleIngestJobSettings {

    private static final long serialVersionUID = 1L;

    private final String endpoint;
    private final int maxFileSize;
    private final List<String> extensions;
    private final Map<String, String> additionalParams;

    public SteganalysisIngestJobSettings(String endpoint, int maxFileSize, List<String> extensions, Map<String, String> additionalParams) {
        this.endpoint = endpoint;
        this.maxFileSize = maxFileSize;
        this.extensions = extensions;
        this.additionalParams = additionalParams;
    }

    public SteganalysisIngestJobSettings() {
        this.endpoint = "localhost:5000";
        this.maxFileSize = 5 * 1024 * 1024;
        this.extensions = new LinkedList<>();
        this.extensions.add("png");
        this.extensions.add("jpg");
        this.additionalParams = new HashMap<String, String>();
    }

    public String getEndpoint() {
        return endpoint;
    }

    public int getMaxFileSize() {
        return maxFileSize;
    }

    public List<String> getExtensions() {
        return extensions;
    }

    public Map<String, String> getAdditionalParams() {
        return additionalParams;
    }

    @Override
    public long getVersionNumber() {
        return serialVersionUID;
    }
}
