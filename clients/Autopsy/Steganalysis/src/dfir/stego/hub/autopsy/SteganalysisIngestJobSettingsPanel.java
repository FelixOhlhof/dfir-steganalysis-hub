/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/GUIForms/JPanel.java to edit this template
 */
package dfir.stego.hub.autopsy;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import javax.swing.table.DefaultTableModel;
import org.sleuthkit.autopsy.ingest.IngestModuleIngestJobSettings;
import org.sleuthkit.autopsy.ingest.IngestModuleIngestJobSettingsPanel;

/**
 *
 * @author felix
 */
public class SteganalysisIngestJobSettingsPanel extends IngestModuleIngestJobSettingsPanel {

    /**
     * Creates new form SteganalysisIngestJobSettingsPanel
     */
    public SteganalysisIngestJobSettingsPanel(SteganalysisIngestJobSettings settings) {
        initComponents();
        customizeComponents(settings);
    }

    private void customizeComponents(SteganalysisIngestJobSettings settings) {
        if (settings != null) {
            endpointTextField.setText(settings.getEndpoint());

            if (settings.getExtensions().contains("png")) {
                pngCheckbox.setState(true);
            }

            if (settings.getExtensions().contains("jpg")) {
                jpgCheckbox.setState(true);
            }

            for (Map.Entry<String, String> entry : settings.getAdditionalParams().entrySet()) {
                ((DefaultTableModel) parameterTable.getModel()).addRow(new String[]{entry.getKey(), entry.getValue()});
            }
        }
    }

    /**
     * This method is called from within the constructor to initialize the form.
     * WARNING: Do NOT modify this code. The content of this method is always
     * regenerated by the Form Editor.
     */
    @SuppressWarnings("unchecked")
    // <editor-fold defaultstate="collapsed" desc="Generated Code">//GEN-BEGIN:initComponents
    private void initComponents() {

        settingsHeadlineLabel = new javax.swing.JLabel();
        endpointTextField = new javax.swing.JTextField();
        endpointLabel = new javax.swing.JLabel();
        fileTypesLable = new javax.swing.JLabel();
        pngCheckbox = new java.awt.Checkbox();
        jpgCheckbox = new java.awt.Checkbox();
        parameterLable = new javax.swing.JLabel();
        jScrollPane1 = new javax.swing.JScrollPane();
        parameterTable = new javax.swing.JTable();
        addRowLable = new javax.swing.JLabel();
        removeRowLable = new javax.swing.JLabel();
        maxFileSizeLable = new java.awt.Label();
        maxFileSizeTextField = new java.awt.TextField();
        mbLable = new javax.swing.JLabel();

        setPreferredSize(new java.awt.Dimension(300, 300));

        settingsHeadlineLabel.setFont(new java.awt.Font("Segoe UI", 1, 14)); // NOI18N
        org.openide.awt.Mnemonics.setLocalizedText(settingsHeadlineLabel, org.openide.util.NbBundle.getMessage(SteganalysisIngestJobSettingsPanel.class, "SteganalysisIngestJobSettingsPanel.settingsHeadlineLabel.text")); // NOI18N

        endpointTextField.setText(org.openide.util.NbBundle.getMessage(SteganalysisIngestJobSettingsPanel.class, "SteganalysisIngestJobSettingsPanel.endpointTextField.text")); // NOI18N

        org.openide.awt.Mnemonics.setLocalizedText(endpointLabel, org.openide.util.NbBundle.getMessage(SteganalysisIngestJobSettingsPanel.class, "SteganalysisIngestJobSettingsPanel.endpointLabel.text")); // NOI18N

        org.openide.awt.Mnemonics.setLocalizedText(fileTypesLable, org.openide.util.NbBundle.getMessage(SteganalysisIngestJobSettingsPanel.class, "SteganalysisIngestJobSettingsPanel.fileTypesLable.text")); // NOI18N

        pngCheckbox.setLabel(org.openide.util.NbBundle.getMessage(SteganalysisIngestJobSettingsPanel.class, "SteganalysisIngestJobSettingsPanel.pngCheckbox.label")); // NOI18N

        jpgCheckbox.setLabel(org.openide.util.NbBundle.getMessage(SteganalysisIngestJobSettingsPanel.class, "SteganalysisIngestJobSettingsPanel.jpgCheckbox.label")); // NOI18N

        org.openide.awt.Mnemonics.setLocalizedText(parameterLable, org.openide.util.NbBundle.getMessage(SteganalysisIngestJobSettingsPanel.class, "SteganalysisIngestJobSettingsPanel.parameterLable.text")); // NOI18N

        parameterTable.setModel(new javax.swing.table.DefaultTableModel(
            new Object [][] {

            },
            new String [] {
                "Name", "Value"
            }
        ) {
            Class[] types = new Class [] {
                java.lang.String.class, java.lang.String.class
            };

            public Class getColumnClass(int columnIndex) {
                return types [columnIndex];
            }
        });
        jScrollPane1.setViewportView(parameterTable);

        addRowLable.setFont(new java.awt.Font("Segoe UI", 1, 24)); // NOI18N
        org.openide.awt.Mnemonics.setLocalizedText(addRowLable, org.openide.util.NbBundle.getMessage(SteganalysisIngestJobSettingsPanel.class, "SteganalysisIngestJobSettingsPanel.addRowLable.text")); // NOI18N
        addRowLable.setHorizontalTextPosition(javax.swing.SwingConstants.CENTER);
        addRowLable.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                addRowLableMouseClicked(evt);
            }
        });

        removeRowLable.setFont(new java.awt.Font("Segoe UI", 1, 24)); // NOI18N
        org.openide.awt.Mnemonics.setLocalizedText(removeRowLable, org.openide.util.NbBundle.getMessage(SteganalysisIngestJobSettingsPanel.class, "SteganalysisIngestJobSettingsPanel.removeRowLable.text")); // NOI18N
        removeRowLable.setHorizontalTextPosition(javax.swing.SwingConstants.CENTER);
        removeRowLable.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                removeRowLableMouseClicked(evt);
            }
        });

        maxFileSizeLable.setText(org.openide.util.NbBundle.getMessage(SteganalysisIngestJobSettingsPanel.class, "SteganalysisIngestJobSettingsPanel.maxFileSizeLable.text")); // NOI18N

        maxFileSizeTextField.setText(org.openide.util.NbBundle.getMessage(SteganalysisIngestJobSettingsPanel.class, "SteganalysisIngestJobSettingsPanel.maxFileSizeTextField.text")); // NOI18N

        org.openide.awt.Mnemonics.setLocalizedText(mbLable, org.openide.util.NbBundle.getMessage(SteganalysisIngestJobSettingsPanel.class, "SteganalysisIngestJobSettingsPanel.mbLable.text")); // NOI18N

        javax.swing.GroupLayout layout = new javax.swing.GroupLayout(this);
        this.setLayout(layout);
        layout.setHorizontalGroup(
            layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(layout.createSequentialGroup()
                .addContainerGap()
                .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING, false)
                    .addComponent(settingsHeadlineLabel, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
                    .addComponent(endpointLabel, javax.swing.GroupLayout.PREFERRED_SIZE, 101, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(jScrollPane1, javax.swing.GroupLayout.PREFERRED_SIZE, 0, Short.MAX_VALUE)
                    .addGroup(layout.createSequentialGroup()
                        .addComponent(parameterLable, javax.swing.GroupLayout.PREFERRED_SIZE, 80, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
                        .addComponent(addRowLable)
                        .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                        .addComponent(removeRowLable, javax.swing.GroupLayout.PREFERRED_SIZE, 14, javax.swing.GroupLayout.PREFERRED_SIZE))
                    .addComponent(endpointTextField)
                    .addGroup(javax.swing.GroupLayout.Alignment.TRAILING, layout.createSequentialGroup()
                        .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                            .addGroup(layout.createSequentialGroup()
                                .addComponent(fileTypesLable, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
                                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED))
                            .addGroup(layout.createSequentialGroup()
                                .addComponent(pngCheckbox, javax.swing.GroupLayout.PREFERRED_SIZE, 60, javax.swing.GroupLayout.PREFERRED_SIZE)
                                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                                .addComponent(jpgCheckbox, javax.swing.GroupLayout.DEFAULT_SIZE, 52, Short.MAX_VALUE)
                                .addGap(15, 15, 15)))
                        .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                            .addComponent(maxFileSizeLable, javax.swing.GroupLayout.PREFERRED_SIZE, 137, javax.swing.GroupLayout.PREFERRED_SIZE)
                            .addGroup(layout.createSequentialGroup()
                                .addComponent(maxFileSizeTextField, javax.swing.GroupLayout.PREFERRED_SIZE, 89, javax.swing.GroupLayout.PREFERRED_SIZE)
                                .addGap(0, 0, 0)
                                .addComponent(mbLable, javax.swing.GroupLayout.PREFERRED_SIZE, 30, javax.swing.GroupLayout.PREFERRED_SIZE)))))
                .addContainerGap(javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE))
        );
        layout.setVerticalGroup(
            layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(layout.createSequentialGroup()
                .addContainerGap()
                .addComponent(settingsHeadlineLabel)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.UNRELATED)
                .addComponent(endpointLabel)
                .addGap(4, 4, 4)
                .addComponent(endpointTextField, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.UNRELATED)
                .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.TRAILING)
                    .addGroup(layout.createSequentialGroup()
                        .addComponent(fileTypesLable)
                        .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                        .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.TRAILING)
                            .addComponent(pngCheckbox, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                            .addComponent(jpgCheckbox, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)))
                    .addGroup(layout.createSequentialGroup()
                        .addComponent(maxFileSizeLable, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addGap(6, 6, 6)
                        .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                            .addComponent(maxFileSizeTextField, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                            .addComponent(mbLable))))
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING, false)
                    .addComponent(addRowLable, javax.swing.GroupLayout.Alignment.TRAILING, javax.swing.GroupLayout.PREFERRED_SIZE, 16, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(removeRowLable, javax.swing.GroupLayout.Alignment.TRAILING, javax.swing.GroupLayout.PREFERRED_SIZE, 16, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(parameterLable))
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(jScrollPane1, javax.swing.GroupLayout.DEFAULT_SIZE, 110, Short.MAX_VALUE)
                .addContainerGap())
        );
    }// </editor-fold>//GEN-END:initComponents

    private void addRowLableMouseClicked(java.awt.event.MouseEvent evt) {//GEN-FIRST:event_addRowLableMouseClicked
        ((DefaultTableModel) parameterTable.getModel()).addRow(new String[]{"", ""});
    }//GEN-LAST:event_addRowLableMouseClicked

    private void removeRowLableMouseClicked(java.awt.event.MouseEvent evt) {//GEN-FIRST:event_removeRowLableMouseClicked
        int selectedRow = parameterTable.getSelectedRow();
        if (selectedRow != -1) {
            ((DefaultTableModel) parameterTable.getModel()).removeRow(selectedRow);
        } else {
            ((DefaultTableModel) parameterTable.getModel()).removeRow(parameterTable.getRowCount() - 1);
        }
    }//GEN-LAST:event_removeRowLableMouseClicked

    private Map<String, String> getKeyValueMap() {
        Map<String, String> map = new HashMap<>();
        for (int row = 0; row < parameterTable.getModel().getRowCount(); row++) {
            String key = (String) parameterTable.getModel().getValueAt(row, 0);
            String value = (String) parameterTable.getModel().getValueAt(row, 1);
            if (!key.isEmpty() && !value.isEmpty()) {
                map.put(key, value);
            }
        }
        return map;
    }

    @Override
    public IngestModuleIngestJobSettings getSettings() {
        String endpoint = endpointTextField.getText().trim();
        List<String> extensions = new LinkedList<>();
        if (pngCheckbox.getState()) {
            extensions.add("png");
        }
        if (jpgCheckbox.getState()) {
            extensions.add("jpg");
        }

        int maxFileSize = -1;

        if (!maxFileSizeTextField.getText().equals("")) {
            maxFileSize = (int) Double.parseDouble(maxFileSizeTextField.getText()) * 1024 * 1024;
        }

        return new SteganalysisIngestJobSettings(endpoint, maxFileSize, extensions, getKeyValueMap());
    }

    // Variables declaration - do not modify//GEN-BEGIN:variables
    private javax.swing.JLabel addRowLable;
    private javax.swing.JLabel endpointLabel;
    private javax.swing.JTextField endpointTextField;
    private javax.swing.JLabel fileTypesLable;
    private javax.swing.JScrollPane jScrollPane1;
    private java.awt.Checkbox jpgCheckbox;
    private java.awt.Label maxFileSizeLable;
    private java.awt.TextField maxFileSizeTextField;
    private javax.swing.JLabel mbLable;
    private javax.swing.JLabel parameterLable;
    private javax.swing.JTable parameterTable;
    private java.awt.Checkbox pngCheckbox;
    private javax.swing.JLabel removeRowLable;
    private javax.swing.JLabel settingsHeadlineLabel;
    // End of variables declaration//GEN-END:variables
}
