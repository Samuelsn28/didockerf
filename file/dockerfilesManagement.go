package file

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"didockerf/model"
)

const (
	defaultPrefix                = `dockerfile_`
	savedDockerfileFileNameRegex = `^` + defaultPrefix + `[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*_[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*$`
	dockerfileIdenfierRegex      = `^[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*:[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*$`
)

var DockerfileIsntCorrectlyFormattedError = errors.New("Dockerfile name or tag isn't correctly formaftted.")

func GetSavedDockerfileFileNameOf(dockerfile model.Dockerfile) (string, error) {
	fileName := defaultPrefix + dockerfile.Name + "_" + dockerfile.Tag

	if !isSavedDockerfileFileNameCorretlyFormatted(fileName) {
		return "", DockerfileIsntCorrectlyFormattedError
	}
	return fileName, nil
}

func getNameOfSavedDockerfile(fileName string) (string, error) {
	if !isSavedDockerfileFileNameCorretlyFormatted(fileName) {
		return "", DockerfileIsntCorrectlyFormattedError
	}
	fileNameSplitted := strings.Split(fileName, "_")

	return fileNameSplitted[1], nil
}

func getTagOfSavedDockerfile(fileName string) (string, error) {
	if !isSavedDockerfileFileNameCorretlyFormatted(fileName) {
		return "", DockerfileIsntCorrectlyFormattedError
	}
	fileNameSplitted := strings.Split(fileName, "_")

	return fileNameSplitted[2], nil
}

func getOriginPathOfSavedDockerfile(fileName string) (string, error) {
	if !isSavedDockerfileFileNameCorretlyFormatted(fileName) {
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

func isSavedDockerfileFileNameCorretlyFormatted(fileName string) bool {
	prepatedRegex, errCompile := regexp.Compile(savedDockerfileFileNameRegex)
	result := prepatedRegex.MatchString(fileName)

	if errCompile != nil {
		return false
	}
	return result
}

func GetSavedDockerfileFileNameOfIdentifer(identifier string) (string, error) {
	identifierSplitted := strings.Split(identifier, ":")

	if len(identifierSplitted) != 2 {
		return "", errors.New("Identifier isn't correctly formatted.")
	}

	name := identifierSplitted[0]
	tag := identifierSplitted[1]

	fileName := defaultPrefix + name + "_" + tag
	return fileName, nil
}

func IsIdentifierCorrectlyFormatted(identifier string) bool {
	preparedRegex, errCompile := regexp.Compile(dockerfileIdenfierRegex)

	if errCompile != nil {
		return false
	}
	return preparedRegex.MatchString(identifier)
}

func TransformSavedDockerfileIdentifierIntoDockerfile(identifier string) (model.Dockerfile, error) {
	savedDockerfileFileName, err := GetSavedDockerfileFileNameOfIdentifer(identifier)
	if err != nil {
		return model.Dockerfile{}, fmt.Errorf("Error: it was not possible to transform identifer into dockerfile: %w", err)
	}

	name, errOnGetName := getNameOfSavedDockerfile(savedDockerfileFileName)

	if errOnGetName != nil {
		return model.Dockerfile{}, fmt.Errorf("Error: it was not possible to get name of saved dockerfile: %w", errOnGetName)
	}

	tag, errOnGetTag := getTagOfSavedDockerfile(savedDockerfileFileName)

	if errOnGetTag != nil {
		return model.Dockerfile{}, fmt.Errorf("Error: it was not possible to get tag of saved dockerfile: %w", errOnGetTag)
	}

	originPath, errOnGetOriginPath := getOriginPathOfSavedDockerfile(savedDockerfileFileName)

	if errOnGetOriginPath != nil {
		return model.Dockerfile{}, fmt.Errorf("Error: it was not possible to get origin path of saved dockerfile: %w", errOnGetOriginPath)
	}

	return model.Dockerfile{
		Name:       name,
		Tag:        tag,
		OriginPath: originPath,
	}, nil
}

func ChangeSavedDockerfileIdentifier(currentIdentifier string, newIdentifier string) error {
	if !IsIdentifierCorrectlyFormatted(currentIdentifier) || !IsIdentifierCorrectlyFormatted(newIdentifier) {
		return errors.New("Identifiers aren't correctly formatted.")
	}

	fileName, errToGetName := GetSavedDockerfileFileNameOfIdentifer(currentIdentifier)

	if errToGetName != nil {
		return errToGetName
	}

	originPath, errToGetOriginPath := getOriginPathOfSavedDockerfile(fileName)

	newFileName, errToGetNewName := GetSavedDockerfileFileNameOfIdentifer(newIdentifier)
	newName, _ := getNameOfSavedDockerfile(newFileName)
	newFileNamePath := defaultSavedDockerfilesDir + "/" + newName + "/" + newFileName

	if errToGetOriginPath != nil {
		return errToGetOriginPath
	}

	if errToGetNewName != nil {
		return errToGetNewName
	}

	currentName, errGettingName := getNameOfSavedDockerfile(fileName)
	if errGettingName != nil {
		return errors.New("Erro ao pegar o nome")
	}

	if currentName == newName {

		fmt.Printf("originPath: %s \n", originPath)
		fmt.Printf("Current name: %s \n", currentName)
		fmt.Printf("O novo nome: %s \n", newName)
		fmt.Printf("O novo path: %s \n", newFileNamePath)

		errRenaming := renameFile(originPath, newFileNamePath)

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

	errCopyingFile := copyFile(originPath, newFileNamePath)
	if errCopyingFile != nil {
		return errors.New("Erro ao copiar o arquivo antigo")
	}

	errDeletingFile := deleteFile(originPath)
	if errDeletingFile != nil {
		return errors.New("Erro ao deletar o save antigo")
	}

	originWithOnlyPath := getPathWithoutName(originPath)
	if isDirEmpty(originWithOnlyPath) {
		errDirIsntEmpty := deleteDir(originWithOnlyPath)

		if errDirIsntEmpty != nil {
			return errors.New("Error: origin dir isn't empty.")
		}
	}

	return nil
}
