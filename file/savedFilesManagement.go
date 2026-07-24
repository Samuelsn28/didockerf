package file

import (
	"fmt"
	"regexp"
	"strings"

	"didockerf/model"
)

const defaultSavedDockerfilesDir = "saves/dockerfiles"

func dockerfilesSaveDirExists() bool {
	return dirExist(defaultSavedDockerfilesDir)
}

func SaveDockerfile(dockerfile model.Dockerfile) {
	if (!existSaveDirOfDockerfile(dockerfile)) {
		createSaveDirOfDockerfile(dockerfile)
	}
	if (existSaveOfDockerfile(dockerfile)) {
		return
	}

	saveDockerfileInItsDir(dockerfile)
}

func existSaveDirOfDockerfile(dockerfile model.Dockerfile) bool {
	saveDirPathOfDockerfile := defaultSavedDockerfilesDir + "/" + dockerfile.Name

	return dirExist(saveDirPathOfDockerfile)
}

func existSaveOfDockerfile(dockerfile model.Dockerfile) bool {
	savedDockerfilePath := defaultSavedDockerfilesDir + "/" + dockerfile.Name + "/" + dockerfile.GetFileName()

	return FileExist(savedDockerfilePath)
}

func createSaveDirOfDockerfile(dockerfile model.Dockerfile) {
	saveDirPathOfDockerfile := defaultSavedDockerfilesDir + "/" + dockerfile.Name
	err := createDir(saveDirPathOfDockerfile)

	if (err != nil) {
		fmt.Println("Erro ao criar dockerfile save dir")
	}
}

func saveDockerfileInItsDir(dockerfile model.Dockerfile) {
	savedDockerfilePath := defaultSavedDockerfilesDir + "/" + dockerfile.Name + "/" + dockerfile.GetFileName()

	err := copyFile(dockerfile.OriginPath, savedDockerfilePath)
	if (err != nil) {
		fmt.Println("Erro ao save dockerfile")
	}
}

func GetAllSavedDockerfiles() []model.Dockerfile {
	savedDockerfilesPaths, err := getAllFilesPathRecursively(defaultSavedDockerfilesDir)

	if (err != nil) {
		fmt.Println("Erro ao get all saved dockerfiles")
	}

	return transformIntoDockerfiles(savedDockerfilesPaths)
}

func transformIntoDockerfiles(savedDockerfilesPaths []string) []model.Dockerfile {
	dockerfiles := []model.Dockerfile{}
	regex, err := regexp.Compile(`^dockerfile_[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*_[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*$`)

	if (err != nil) {
		fmt.Println("Regex inválida")
		return []model.Dockerfile{}
	}

	for _, savedDockerfilePath := range(savedDockerfilesPaths) {
		savedDockerfilePathSplitted := strings.Split(savedDockerfilePath, "/")
		savedDockerfileFileName := savedDockerfilePathSplitted[ len(savedDockerfilePathSplitted) - 1 ]

		if (!regex.MatchString(savedDockerfileFileName)) {
			continue
		}
		savedDockerfileFileNameSplitted := strings.Split(savedDockerfileFileName, "_")

		dockerfiles = append(dockerfiles, model.Dockerfile{
			Name: savedDockerfileFileNameSplitted[1],
			Version: savedDockerfileFileNameSplitted[2],
			OriginPath: savedDockerfilePath,
		})
	}
	return dockerfiles
}
