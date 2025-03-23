/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package dfir.stego.hub.autopsy;

import org.sleuthkit.datamodel.BlackboardAttribute;
import org.sleuthkit.datamodel.SleuthkitCase;
import org.sleuthkit.autopsy.casemodule.Case;
import org.sleuthkit.datamodel.BlackboardArtifact;
import org.sleuthkit.datamodel.TskCoreException;

public class CustomAttributeHelper {

    private static BlackboardAttribute.Type customType;

    /**
     * Registriert einen benutzerdefinierten BlackboardAttribute-Typen.
     */
    public static BlackboardAttribute.Type createCustomAttributeType(String typeName, String displayName) throws TskCoreException {
        SleuthkitCase sleuthkitCase = Case.getCurrentCase().getSleuthkitCase();

        // Prüfen, ob der Attributtyp bereits existiert
        BlackboardAttribute.Type existingType = sleuthkitCase.getAttributeType(typeName);
        if (existingType != null) {
            return existingType;
        }

        // Wenn der Typ nicht existiert, neuen Attribut-Typ hinzufügen
        sleuthkitCase.addAttrType(typeName, displayName);

        // Nach dem Erstellen den Typ abrufen und zurückgeben
        return sleuthkitCase.getAttributeType(typeName);
    }

    /**
     * Beispielmethode zum Erstellen eines benutzerdefinierten Attributs und
     * Speichern des Attributs in einem Artefakt.
     */
    public static void addCustomAttributeToArtifact(BlackboardArtifact artifact, String typeName, String displayName, String value) {
        try {
            // Erstelle oder hole den benutzerdefinierten Attribut-Typen
            customType = createCustomAttributeType(typeName, displayName);

            // Erstelle das benutzerdefinierte Attribut
            BlackboardAttribute customAttribute = new BlackboardAttribute(customType, SteganalysisModuleFactory.getModuleName(), value);

            // Füge das Attribut dem Artefakt hinzu
            artifact.addAttribute(customAttribute);

        } catch (TskCoreException e) {
            System.err.println("error while adding attribute to artifact: " + e.getMessage());
        }
    }
}
