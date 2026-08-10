package savesManagement

import (
	"errors"

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

func CopySaveComposeFileTo(identifierStr string, destination string) error {
	composeFileType := fy.GetComposeFileType()
	savedComposeFileIdentifier, errOnGetIdentifier := model.CreateIdentifierFromStr(identifierStr, composeFileType)

	if errOnGetIdentifier != nil {
		return errOnGetIdentifier
	}

	return copySavedFileTo(savedComposeFileIdentifier, destination)
}

func ChangeSavedComposeFileIdentifier(savedIdentifierStr string, newIdentifierStr string) error {
	composeFileType := fy.GetComposeFileType()
	savedIdentifier, errOnGetSavedIdentifier := model.CreateIdentifierFromStr(savedIdentifierStr, composeFileType)
	newIdentifier, errOnGetNewIdentifier := model.CreateIdentifierFromStr(newIdentifierStr, composeFileType)

	if errOnGetSavedIdentifier != nil {
		return errors.Join(errors.New("Saved compose file identifier:"), errOnGetSavedIdentifier)
	}
	if errOnGetNewIdentifier != nil {
		return errors.Join(errors.New("New identifier:"), errOnGetNewIdentifier)
	}

	return changeIdentifierOfSavedFile(savedIdentifier, newIdentifier)
}

func RemoveSavedComposeFile(savedIdentifierStr string) error {
	composeFileType := fy.GetComposeFileType()
	savedIdentifier, errOnGetIdentifier := model.CreateIdentifierFromStr(savedIdentifierStr, composeFileType)

	if errOnGetIdentifier != nil {
		return errOnGetIdentifier
	}

	return removeSavedFile(savedIdentifier)
}

func RemoveAllSavedComposeFiles() error {
	return removeAllSavedFiles(fy.GetComposeFileType())
}
