package file

import (
	"errors"
	"fmt"
	"strings"

	"didockerf/model"
	"didockerf/out"
	"didockerf/util"
)

const (
	defaultSavedDockerfilesDir = "saves/dockerfiles"
	separetorSavedDockerfile   = `_`
	defaultPrefix              = `dockerfile` + separetorSavedDockerfile

	savedDockerfileFileNameRegex = `^` + defaultPrefix + `[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*` + separetorSavedDockerfile + `[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*$`
)

var DockerfileIsntCorrectlyFormattedError = errors.New("Dockerfile name or tag isn't correctly formaftted.")

func GetSavedDockerfileFileNameOfIdentiferStr(identifierStr model.IdentifierStr) (string, error) {
	if !model.IsIdentifierStrValid(identifierStr) {
		return "", errors.New("Identifier isn't correctly formatted.")
	}

	name := identifierStr.GetName()
	tag := identifierStr.GetTag()

	fileName := defaultPrefix + name + separetorSavedDockerfile + tag
	return fileName, nil
}

func TransformSavedDockerfileIdentifierStrIntoDockerfile(identifierStr model.IdentifierStr) (model.Dockerfile, error) {
	savedDockerfileFileName, err := GetSavedDockerfileFileNameOfIdentiferStr(identifierStr)
	if err != nil {
		return model.Dockerfile{}, fmt.Errorf("Error: it was not possible to transform identifer into dockerfile: %w", err)
	}

	savedDockerfileIdentifier := identifierStr.GetIdentifier()
	originPath, errOnGetOriginPath := getOriginPathOfSavedDockerfile(savedDockerfileFileName)

	if errOnGetOriginPath != nil {
		return model.Dockerfile{}, fmt.Errorf("Error: it was not possible to get origin path of saved dockerfile: %w", errOnGetOriginPath)
	}

	newIdentifier := savedDockerfileIdentifier

	return model.CreateDockerfile(
		newIdentifier,
		originPath,
	), nil
}

func ChangeSavedDockerfileIdentifier(currentIdentifierStr model.IdentifierStr, newIdentifierStr model.IdentifierStr) error {
	if !model.IsIdentifierStrValid(currentIdentifierStr) || !model.IsIdentifierStrValid(newIdentifierStr) {
		return errors.New("Identifiers aren't correctly formatted.")
	}

	currentFileName, errToGetName := GetSavedDockerfileFileNameOfIdentiferStr(currentIdentifierStr)
	newFileName, errToGetNewFileName := GetSavedDockerfileFileNameOfIdentiferStr(newIdentifierStr)

	if errToGetName != nil {
		return errToGetName
	}
	if errToGetNewFileName != nil {
		return errToGetNewFileName
	}

	currentName := currentIdentifierStr.GetName()
	newName := newIdentifierStr.GetName()

	currentOriginPath, errToGetOriginPath := getOriginPathOfSavedDockerfile(currentFileName)
	newOriginPath := defaultSavedDockerfilesDir + "/" + newName + "/" + newFileName

	if errToGetOriginPath != nil {
		return errToGetOriginPath
	}

	if currentName == newName {
		errRenaming := renameFile(currentOriginPath, newOriginPath)

		if errRenaming != nil {
			return errRenaming
		}
		return nil
	}
	newDir := defaultSavedDockerfilesDir + "/" + newName

	errCreatingDir := createDir(newDir)
	if errCreatingDir != nil {
		return errors.New("Erro ao criar o dir da mudança")
	}

	errCopyingFile := copyFile(currentOriginPath, newOriginPath)
	if errCopyingFile != nil {
		return errors.New("Erro ao copiar o arquivo antigo")
	}

	errDeletingFile := deleteFile(currentOriginPath)
	if errDeletingFile != nil {
		return errors.New("Erro ao deletar o save antigo")
	}

	currentOriginWithOnlyPath := getPathWithoutName(currentOriginPath)
	if isDirEmpty(currentOriginWithOnlyPath) {
		errDirIsntEmpty := deleteDir(currentOriginWithOnlyPath)

		if errDirIsntEmpty != nil {
			return errors.New("Error: origin dir isn't empty.")
		}
	}

	return nil
}

func SaveDockerfile(dockerfile model.Dockerfile) {
	if !existSaveDirOfDockerfile(dockerfile) {
		errOnCreateSaveDir := createSaveDirOfDockerfile(dockerfile)

		if errOnCreateSaveDir != nil {
			out.FatalError(errOnCreateSaveDir)
			return
		}
	}
	if existSaveOfDockerfile(dockerfile) {
		out.Warn("Already exist dockerfile with the passed identifier.")
		return
	}

	errOnSaveDockerfile := saveDockerfileInItsDir(dockerfile)
	if errOnSaveDockerfile != nil {
		out.FatalError(errOnSaveDockerfile)
	}
}

func existSaveDirOfDockerfile(dockerfile model.Dockerfile) bool {
	saveDirPathOfDockerfile := getSaveDirOfDockerfile(dockerfile)

	return dirExist(saveDirPathOfDockerfile)
}

func createSaveDirOfDockerfile(dockerfile model.Dockerfile) error {
	saveDirPathOfDockerfile := getSaveDirOfDockerfile(dockerfile)
	err := createDir(saveDirPathOfDockerfile)
	if err != nil {
		return errors.New("Error: it was not possible to create the specific save dir for the dockerfile.")
	}

	return nil
}

func existSaveOfDockerfile(dockerfile model.Dockerfile) bool {
	savedDockerfilePath, errToGetPath := getSavedDockerfilePathWithFileNameOf(dockerfile)
	if errToGetPath != nil {
		return false
	}

	return FileExist(savedDockerfilePath)
}

func saveDockerfileInItsDir(dockerfile model.Dockerfile) error {
	destinationPath, errOnGetDestination := getSavedDockerfilePathWithFileNameOf(dockerfile)

	if errOnGetDestination != nil {
		return errOnGetDestination
	}

	errOnSaveDockerfile := copyFile(dockerfile.OriginPath, destinationPath)
	if errOnSaveDockerfile != nil {
		return errors.New("Error: it was not possible to save the dockerfile in the save folder.")
	}
	return nil
}

func getSaveDirOfDockerfile(dockerfile model.Dockerfile) string {
	return defaultSavedDockerfilesDir + "/" + dockerfile.Identifier.Name
}

func GetAllSavedDockerfiles() []model.Dockerfile {
	savedDockerfilesPaths, err := getAllFilesPathRecursively(defaultSavedDockerfilesDir)
	if err != nil {
		out.PrintFatalError("Error: it was not possible read saved dockerfiles in the save folder.")
		return []model.Dockerfile{}
	}

	return transformIntoDockerfiles(savedDockerfilesPaths)
}

func transformIntoDockerfiles(savedDockerfilesPaths []string) []model.Dockerfile {
	dockerfiles := []model.Dockerfile{}

	for _, savedDockerfilePath := range savedDockerfilesPaths {
		savedDockerfileFileName := getNameWithoutPath(savedDockerfilePath)

		name, errOnName := getNameOfSavedDockerfile(savedDockerfileFileName)
		tag, errOnVersion := getTagOfSavedDockerfile(savedDockerfileFileName)

		if errOnName != nil || errOnVersion != nil {
			continue
		}

		newIdentifier, _ := model.CreateIdentifier(name, tag)
		dockerfiles = append(dockerfiles, model.CreateDockerfile(
			newIdentifier,
			savedDockerfilePath,
		))
	}
	return dockerfiles
}

func getSavedDockerfilePathWithFileNameOf(dockerfile model.Dockerfile) (string, error) {
	savedDockerfileFileName, errOnGetFileName := GetSavedDockerfileFileNameOf(dockerfile)

	if errOnGetFileName != nil {
		return "", errOnGetFileName
	}

	return getSaveDirOfDockerfile(dockerfile) + "/" + savedDockerfileFileName, nil
}

func GetSavedDockerfileFileNameOf(dockerfile model.Dockerfile) (string, error) {
	fileName := defaultPrefix + dockerfile.Identifier.Name + separetorSavedDockerfile + dockerfile.Identifier.Tag

	if !IsSavedDockerfileFileNameValid(fileName) {
		return "", DockerfileIsntCorrectlyFormattedError
	}
	return fileName, nil
}

func getNameOfSavedDockerfile(fileName string) (string, error) {
	if !IsSavedDockerfileFileNameValid(fileName) {
		return "", DockerfileIsntCorrectlyFormattedError
	}
	fileNameSplitted := strings.Split(fileName, separetorSavedDockerfile)

	return fileNameSplitted[1], nil
}

func getTagOfSavedDockerfile(fileName string) (string, error) {
	if !IsSavedDockerfileFileNameValid(fileName) {
		return "", DockerfileIsntCorrectlyFormattedError
	}
	fileNameSplitted := strings.Split(fileName, separetorSavedDockerfile)

	return fileNameSplitted[2], nil
}

func getOriginPathOfSavedDockerfile(fileName string) (string, error) {
	if !IsSavedDockerfileFileNameValid(fileName) {
		return "", DockerfileIsntCorrectlyFormattedError
	}

	savedDockerfileName, errToGetName := getNameOfSavedDockerfile(fileName)

	if errToGetName != nil {
		return "", errToGetName
	}

	fullPathSavedDockerfile := defaultSavedDockerfilesDir + "/" + savedDockerfileName + "/" + fileName
	if !FileExist(fullPathSavedDockerfile) {
		return "", errors.New("Saved dockerfile doesn't exist.")
	}
	return fullPathSavedDockerfile, nil
}

func IsSavedDockerfileFileNameValid(fileName string) bool {
	preparedRegex := util.CreatePreparedRegex(savedDockerfileFileNameRegex)

	return preparedRegex.MatchString(fileName)
}

func dockerfilesSaveDirExists() bool {
	return dirExist(defaultSavedDockerfilesDir)
}
