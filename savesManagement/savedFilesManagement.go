package savesManagement

import (
	"errors"
	"strings"

	"didockerf/out"
	"didockerf/savesManagement/internal/model"
	fy "didockerf/savesManagement/internal/model/fileType"
)

const (
	SavesDir = "saves"
)

func copySavedFileTo(savedFileIdentifier model.Identifier, destinationPath string) error {
	if !DirExist(destinationPath) {
		return errors.New("Destination path doesn't exist.")
	}
	if !alreadyExistSaveWithIdentifier(savedFileIdentifier) {
		return errors.New("Don't exist saved file with the passed identifier.")
	}

	savedFileDestination := destinationPath
	if !strings.HasSuffix(savedFileDestination, "/") {
		savedFileDestination = savedFileDestination + "/"
	}
	savedFileDestination = savedFileDestination + getSaveFileNameOfIdentifier(savedFileIdentifier)
	savedFileOrigin := getSavePathOfIdentifier(savedFileIdentifier)

	errOnCopySavedFile := copyFile(savedFileOrigin, savedFileDestination)
	if errOnCopySavedFile != nil {
		return errOnCopySavedFile
	}

	return nil
}

func getAllSavedFilesOf(fileType fy.FileType) []SavedFileInfo {
	dirOfTheFileType := getPathOfFileType(fileType)
	files, errOnGetAllFiles := getAllFilesPathRecursively(dirOfTheFileType)

	if !DirExist(dirOfTheFileType) {
		errOnCreateSaveDir := createSaveDirOfFileType(fileType)
		if errOnCreateSaveDir != nil {
			out.PrintFatalError("Error: it was not possible to create dockerfiles's save dir.")
		}
	} else if errOnGetAllFiles != nil {
		out.PrintFatalError("Error: it was not possible read saved dockerfiles in the save folder.")
	}

	return transformIntoSavedFileInfos(files)
}

func getPathOfFileType(fileType fy.FileType) string {
	return SavesDir + "/" + fileType.GetSaveLocation()
}

func createSaveDirOfFileType(fileType fy.FileType) error {
	return createDir(getPathOfFileType(fileType))
}

func transformIntoSavedFileInfos(filesWithPath []string) []SavedFileInfo {
	savedFileInfos := []SavedFileInfo{}

	for _, fileWithPath := range filesWithPath {
		fileName := getNameWithoutPath(fileWithPath)

		savedFileIdentifier, errOnFileName := model.TransformFileNameIntoIdentifier(fileName)
		if errOnFileName != nil {
			continue
		}

		savedFile := model.CreateSavedFile(
			savedFileIdentifier, fileWithPath,
		)

		savedFileInfos = append(savedFileInfos, createSavedFileInfoOf(savedFile))
	}
	return savedFileInfos
}

func saveFile(identifier model.Identifier, filePath string) error {
	if !alreadyExistSaveDirOfIdentifier(identifier) {
		errOnCreateSaveDir := createSaveDirOfIdentifier(identifier)

		if errOnCreateSaveDir != nil {
			return errOnCreateSaveDir
		}
	}
	if alreadyExistSaveWithIdentifier(identifier) {
		out.Warn("Already exist dockerfile with the passed identifier.")
		return nil
	}

	fileToSave := model.CreateSavedFile(identifier, filePath)
	errOnSaveFile := saveFileInItsDir(fileToSave)
	if errOnSaveFile != nil {
		return errOnSaveFile
	}

	return nil
}

func getSaveFileNameOfIdentifier(identifier model.Identifier) string {
	return model.TransformIdentifierIntoFileName(identifier)
}

func getSaveDirPathOfIdentifier(identifier model.Identifier) string {
	return getPathOfFileType(identifier.FileType) + "/" + identifier.Name
}

func getSavePathOfIdentifier(identifier model.Identifier) string {
	return getSaveDirPathOfIdentifier(identifier) + "/" + getSaveFileNameOfIdentifier(identifier)
}

func alreadyExistSaveDirOfIdentifier(identifier model.Identifier) bool {
	saveDirPath := getSaveDirPathOfIdentifier(identifier)

	return DirExist(saveDirPath)
}

func createSaveDirOfIdentifier(identifier model.Identifier) error {
	saveDirPath := getSaveDirPathOfIdentifier(identifier)

	errOnCreateSaveDir := createDir(saveDirPath)
	if errOnCreateSaveDir != nil {
		return errors.New("Error: it was not possible to create the specific save dir for the identifier.")
	}
	return nil
}

func alreadyExistSaveWithIdentifier(identifier model.Identifier) bool {
	savePath := getSavePathOfIdentifier(identifier)

	return FileExist(savePath)
}

func saveFileInItsDir(fileToSave model.SavedFile) error {
	destinationPath := getSavePathOfIdentifier(fileToSave.Identifier)

	errOnCopyFileToItsDir := copyFile(fileToSave.OriginPath, destinationPath)
	if errOnCopyFileToItsDir != nil {
		return errors.New("Error: it was not possible to save the file in its save dir.")
	}

	return nil
}

func changeIdentifierOfSavedFile(savedFileIdentifier model.Identifier, newIdentifier model.Identifier) error {
	savedFileOriginPath := getSavePathOfIdentifier(savedFileIdentifier)
	newOriginPath := getSavePathOfIdentifier(newIdentifier)

	if savedFileIdentifier.Name == newIdentifier.Name {
		errOnRename := renameFile(savedFileOriginPath, newOriginPath)
		if errOnRename != nil {
			return errOnRename
		}
		return nil
	}

	errOnCreateSaveDir := createSaveDirOfIdentifier(newIdentifier)
	if errOnCreateSaveDir != nil {
		return errOnCreateSaveDir
	}

	errOnCopySave := copyFile(savedFileOriginPath, newOriginPath)
	if errOnCopySave != nil {
		return errOnCopySave
	}

	errOnDeleteOldSave := deleteFile(savedFileOriginPath)
	if errOnDeleteOldSave != nil {
		return errOnDeleteOldSave
	}

	errOnDeleteDirOfOldSave := deleteDirOfOldSaveIfItIsEmpty(savedFileIdentifier)
	if errOnDeleteDirOfOldSave != nil {
		return errOnDeleteDirOfOldSave
	}

	return nil
}

func deleteDirOfOldSaveIfItIsEmpty(oldSaveIdentifier model.Identifier) error {
	dirOfOldSave := getSaveDirPathOfIdentifier(oldSaveIdentifier)
	if !isDirEmpty(dirOfOldSave) {
		return nil
	}

	errOnDeleteDir := deleteDir(dirOfOldSave)

	if errOnDeleteDir != nil {
		return errOnDeleteDir
	}

	return nil
}

func removeSavedFile(savedFileIdentifier model.Identifier) error {
	savedFilePath := getSavePathOfIdentifier(savedFileIdentifier)

	errOnDeleteFile := deleteFile(savedFilePath)
	if errOnDeleteFile != nil {
		return errOnDeleteFile
	}

	errOnDeleteODirOfOldSave := deleteDirOfOldSaveIfItIsEmpty(savedFileIdentifier)
	if errOnDeleteODirOfOldSave != nil {
		return errOnDeleteODirOfOldSave
	}

	return nil
}

func removeAllSavedFiles(fileType fy.FileType) error {
	pathOfTheFileType := getPathOfFileType(fileType)

	errOnDeleteDir := deleteDirWithAllContent(pathOfTheFileType)
	if errOnDeleteDir != nil {
		return errOnDeleteDir
	}

	errOnRecreateDir := createDir(pathOfTheFileType)
	if errOnRecreateDir != nil {
		return errOnRecreateDir
	}

	return nil
}
