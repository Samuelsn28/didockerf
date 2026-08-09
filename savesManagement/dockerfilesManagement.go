package savesManagement

import (
	"errors"

	"didockerf/savesManagement/internal/model"
	fy "didockerf/savesManagement/internal/model/fileType"
)

func GetAllSavedDockerfilesInfos() []SavedFileInfo {
	return getAllSavedFilesOf(fy.GetDockerfileType())
}

func SaveDockerfile(identifierStr string, originPath string) error {
	identifier, errOnGetIdentifier := model.CreateIdentifierFromStr(identifierStr, fy.GetDockerfileType())
	if errOnGetIdentifier != nil {
		return errOnGetIdentifier
	}

	return saveFile(identifier, originPath)
}

func ChangeSavedDockerfileIdentifier(savedDockerfileIdentifierStr string, newIdentifierStr string) error {
	dockerfileType := fy.GetDockerfileType()
	savedDockerfileIdentifier, errOnGetSavedIdentifier := model.CreateIdentifierFromStr(savedDockerfileIdentifierStr, dockerfileType)
	newIdentifier, errOnGetNewIdentifier := model.CreateIdentifierFromStr(newIdentifierStr, dockerfileType)

	if errOnGetSavedIdentifier != nil {
		return errors.Join(errors.New("Saved dockerfile identifier:"), errOnGetSavedIdentifier)
	}
	if errOnGetNewIdentifier != nil {
		return errors.Join(errors.New("New identifier:"), errOnGetNewIdentifier)
	}

	return changeIdentifierOfSavedFile(savedDockerfileIdentifier, newIdentifier)
}

func CopySavedDockerfileTo(savedDockerfileIdentifierStr string, destination string) error {
	dockerfileType := fy.GetDockerfileType()
	savedDockerfileIdentifier, errOnGetSavedIdentifier := model.CreateIdentifierFromStr(savedDockerfileIdentifierStr, dockerfileType)

	if errOnGetSavedIdentifier != nil {
		return errOnGetSavedIdentifier
	}

	return copySavedFileTo(savedDockerfileIdentifier, destination)
}

func RemoveSavedDockerfile(savedDockerfileIdentifierStr string) error {
	dockerfileType := fy.GetDockerfileType()
	savedDockerfileIdentifier, errOnGetSavedIdentifier := model.CreateIdentifierFromStr(savedDockerfileIdentifierStr, dockerfileType)

	if errOnGetSavedIdentifier != nil {
		return errOnGetSavedIdentifier
	}

	return removeSavedFile(savedDockerfileIdentifier)
}
