package model

import (
	"errors"
	"strings"

	fy "didockerf/savesManagement/internal/model/fileType"
	"didockerf/util"
)

const (
	replaceMark        = "_-REPLACE-_"
	separetorSavedFile = "_"

	savedDockerfileFileNameRegex = `^` + replaceMark + separetorSavedFile + `[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*` + separetorSavedFile + `[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*$`
)

type SavedFile struct {
	Identifier Identifier
	OriginPath string
}

func CreateSavedFile(identifier Identifier, originPath string) SavedFile {
	return SavedFile{
		Identifier: identifier,
		OriginPath: originPath,
	}
}

func TransformFileNameIntoIdentifier(fileName string) (Identifier, error) {
	if !IsSavedFileNameValid(fileName) {
		return Identifier{}, errors.New("The file name for the saved file is invalid.")
	}
	fileNameSplitted := strings.Split(fileName, separetorSavedFile)
	fileType := fy.GetFileTypeOfFilePrefix(fileNameSplitted[0])

	if fileType == nil {
		return Identifier{}, errors.New("The file type passed to identifier is invalid.")
	}

	name := fileNameSplitted[1]
	tag := fileNameSplitted[2]

	identifier, _ := CreateIdentifier(name, tag, *fileType)
	return identifier, nil
}

func TransformIdentifierIntoFileName(identifier Identifier) string {
	return identifier.FileType.GetFilePrefix() + separetorSavedFile + identifier.Name + separetorSavedFile + identifier.Tag
}

func IsSavedFileNameValid(savedFileName string) bool {
	avaliableFileTypes := fy.GetAllFileTypes()

	for _, fileType := range avaliableFileTypes {
		preparedRegex := util.CreatePreparedRegexWithReplace(savedDockerfileFileNameRegex, replaceMark, fileType.GetFilePrefix())

		if preparedRegex.MatchString(savedFileName) {
			return true
		}
	}
	return false
}
