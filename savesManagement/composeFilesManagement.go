package savesManagement

import (
	"didockerf/savesManagement/internal/model"
	fy "didockerf/savesManagement/internal/model/fileType"
)

func GetAllSavedComposeFilesInfos() []SavedFileInfo {
	return getAllSavedFilesOf(fy.GetComposeFileType())
}

func SaveComposeFile(identifierStr string, originPath string) error {
	composeFileType := fy.GetComposeFileType()
	identifier, errOnGetIdentifier := model.CreateIdentifierFromStr(identifierStr, composeFileType)
	if errOnGetIdentifier != nil {
		return errOnGetIdentifier
	}

	return saveFile(identifier, originPath)
}
