package file

import (
	"fmt"
)

const saveDir = "saves/dockerfiles"
const savedDockerfileDefaultName = "dockerfile-"

func CreateSaveDirIfNotExists() error {
	if (!saveDirExists()) {
		return createSaveDir()
	}
	return nil
}

func saveDirExists() bool {
	return dirExist(saveDir)
}

func createSaveDir() error {
	err := createDir(saveDir)

	if err != nil {
		fmt.Println("[!] Error creating save dir")
		return err
	}
}

func CreateDockerfileSave(dockerfilePath string, name string, version string) {
	if (!existDockerfileSaveDir(name)) {
		createDockerfileSaveDir(name)
	}
	if (existDockerfileSaveWithVersion(name, version)) {
		return
	}

	saveDockerfile(dockerfilePath, name, version)
}

func existDockerfileSaveDir(saveName string) bool {
	dockerfileSaveDirPath := saveDir + "/" + saveName

	return dirExist(dockerfileSaveDirPath)
}

func existDockerfileSaveWithVersion(name string, version string) bool {
	dockerfileSaveName := savedDockerfileDefaultName + name + "-" + version
	dockerfileSavePath := saveDir + "/" + name + "/" + dockerfileSaveName

	return FileExist(dockerfileSavePath)
}

func createDockerfileSaveDir(saveName string) {
	dockerfileSaveDirPath := saveDir + "/" + saveName
	err := createDir(dockerfileSaveDirPath)

	if err != nil {
		fmt.Println("Erro ao criar dockerfile save dir")
	}
}

func saveDockerfile(dockerfilePath string, name string, version string) {
	dockerfileSaveName := savedDockerfileDefaultName + name + "-" + version
	dockerfileSavePath := saveDir + "/" + name + "/" + dockerfileSaveName

	err := copyFile(dockerfilePath, dockerfileSavePath)
	if err != nil {
		fmt.Println("Erro ao save dockerfile")
	}
}
