/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package dfir.stego.hub.autopsy;

import org.openide.util.NbBundle;
import org.sleuthkit.datamodel.BlackboardArtifact;
import org.sleuthkit.datamodel.SleuthkitCase;
import org.sleuthkit.autopsy.casemodule.Case;
import org.sleuthkit.datamodel.AbstractFile;
import org.sleuthkit.datamodel.TskCoreException;
import org.sleuthkit.datamodel.TskDataException;

public class CustomArtifactHelper {

    private static BlackboardArtifact.Type customArtifactType = null;

    public static BlackboardArtifact.Type createCustomArtifactType(String typeName, String displayName) throws TskCoreException, TskDataException {
        SleuthkitCase sleuthkitCase = Case.getCurrentCase().getSleuthkitCase();

        BlackboardArtifact.Type existingType = sleuthkitCase.getArtifactType(displayName);
        if (existingType != null) {
            return existingType;
        }

        BlackboardArtifact.Type newType = sleuthkitCase.addBlackboardArtifactType(typeName, displayName);

        return newType;
    }

    public static BlackboardArtifact addCustomArtifactToFile(AbstractFile file) throws TskDataException, TskCoreException {

        customArtifactType = createCustomArtifactType("TSK_STEGANALYSIS", NbBundle.getMessage(SteganalysisModuleFactory.class, "moduleDisplayName.text"));

        BlackboardArtifact artifact = file.newArtifact(customArtifactType.getTypeID());

        return artifact;
    }
}
