package file

import (
	"errors"
	"regexp"
	"strings"

	"didockerf/model"
)

const defaultPrefix = "dockerfile_"
const savedDockerfileFileNameRegex = `^` + defaultPrefix + `[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*_[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*$`

var DockerfileIsntCorrectlyFormattedError = errors.New("Dockerfile name or tag isn't correctly formatted.")

func GetSavedDockerfileFileNameOf(dockerfile model.Dockerfile) (string, error) {
	fileName := defaultPrefix + dockerfile.Name + "_" + dockerfile.Tag

	if (!isSavedDockerfileFileNameCorretlyFormatted(fileName)) {
		return "", DockerfileIsntCorrectlyFormattedError
	}
	return fileName, nil
}

func getNameOfSavedDockerfile(fileName string) (string, error) {
	if (!isSavedDockerfileFileNameCorretlyFormatted(fileName)) {
		return "", DockerfileIsntCorrectlyFormattedError
	}
	fileNameSplitted := strings.Split(fileName, "_") 

	return fileNameSplitted[1], nil
}

func getTagOfSavedDockerfile(fileName string) (string, error) {
	if (!isSavedDockerfileFileNameCorretlyFormatted(fileName)) {
		return "", DockerfileIsntCorrectlyFormattedError
	}
	fileNameSplitted := strings.Split(fileName, "_") 

	return fileNameSplitted[2], nil
}

func isSavedDockerfileFileNameCorretlyFormatted(fileName string) bool {
	prepatedRegex, errCompile := regexp.Compile(savedDockerfileFileNameRegex)
	result := prepatedRegex.MatchString(fileName)

	if (errCompile != nil) {
		return false
	}
	return result
}

