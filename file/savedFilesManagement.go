package file

import (
	"fmt"

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

	if err != nil {
		fmt.Println("Erro ao criar dockerfile save dir")
	}
}

func saveDockerfileInItsDir(dockerfile model.Dockerfile) {
	savedDockerfilePath := defaultSavedDockerfilesDir + "/" + dockerfile.Name + "/" + dockerfile.GetFileName()

	err := copyFile(dockerfile.OriginPath, savedDockerfilePath)
	if err != nil {
		fmt.Println("Erro ao save dockerfile")
	}
}
