package file

import (
	"errors"

	"didockerf/model"
	"didockerf/out"
)

const defaultSavedDockerfilesDir = "saves/dockerfiles"

func dockerfilesSaveDirExists() bool {
	return dirExist(defaultSavedDockerfilesDir)
}

func SaveDockerfile(dockerfile model.Dockerfile) {
	if (!existSaveDirOfDockerfile(dockerfile)) {
		errOnCreateSaveDir := createSaveDirOfDockerfile(dockerfile)

		if (errOnCreateSaveDir != nil){
			out.PrintError(errOnCreateSaveDir)
			return
		}
	}
	if (existSaveOfDockerfile(dockerfile)) {
		out.PrintWarn("Already exist dockerfile with the passed identifier.")
		return
	}

	errOnSaveDockerfile := saveDockerfileInItsDir(dockerfile)
	if (errOnSaveDockerfile != nil) {
		out.PrintError(errOnSaveDockerfile)
	}
}

func existSaveDirOfDockerfile(dockerfile model.Dockerfile) bool {
	saveDirPathOfDockerfile := defaultSavedDockerfilesDir + "/" + dockerfile.Name

	return dirExist(saveDirPathOfDockerfile)
}

func existSaveOfDockerfile(dockerfile model.Dockerfile) bool {
	savedDockerfileFileName, err := GetSavedDockerfileFileNameOf(dockerfile)

	if (err != nil) {
		out.PrintError(err)
		return false
	}

	savedDockerfilePath := defaultSavedDockerfilesDir + "/" + dockerfile.Name + "/" + savedDockerfileFileName

	return FileExist(savedDockerfilePath)
}

func createSaveDirOfDockerfile(dockerfile model.Dockerfile) error {
	saveDirPathOfDockerfile := defaultSavedDockerfilesDir + "/" + dockerfile.Name
	err := createDir(saveDirPathOfDockerfile)

	if (err != nil) {
		return errors.New("Error: it was not possible to create the save folder for dockerfiles.")
	}
	return nil
}

func saveDockerfileInItsDir(dockerfile model.Dockerfile) error {
	savedDockerfileFileName, errOnGetFileName := GetSavedDockerfileFileNameOf(dockerfile)

	if (errOnGetFileName != nil) {
		return errOnGetFileName
	}
	savedDockerfilePath := defaultSavedDockerfilesDir + "/" + dockerfile.Name + "/" + savedDockerfileFileName

	errOnSaveDockerfile := copyFile(dockerfile.OriginPath, savedDockerfilePath)
	if (errOnSaveDockerfile != nil) {
		return errors.New("Error: it was not possible to save the dockerfile in the save folder.")
	}
	return nil
}

func GetAllSavedDockerfiles() []model.Dockerfile {
	savedDockerfilesPaths, err := getAllFilesPathRecursively(defaultSavedDockerfilesDir)

	if (err != nil) {
		out.PrintError(errors.New("Error: it was not possible read saved dockerfiles in the save folder."))
		return []model.Dockerfile{}
	}

	return transformIntoDockerfiles(savedDockerfilesPaths)
}

func transformIntoDockerfiles(savedDockerfilesPaths []string) []model.Dockerfile {
	dockerfiles := []model.Dockerfile{}

	for _, savedDockerfilePath := range(savedDockerfilesPaths) {
		savedDockerfileFileName := getNameWithoutPath(savedDockerfilePath)

		name, errOnName := getNameOfSavedDockerfile(savedDockerfileFileName)
		tag, errOnVersion := getTagOfSavedDockerfile(savedDockerfileFileName)

		if (errOnName != nil || errOnVersion != nil) {
			continue
		}

		dockerfiles = append(dockerfiles, model.Dockerfile{
			Name: name,
			Tag: tag,
			OriginPath: savedDockerfilePath,
		})
	}
	return dockerfiles
}
