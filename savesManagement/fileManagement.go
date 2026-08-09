package savesManagement

import (
	"io"
	"os"
	"path/filepath"
)

func FileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func createFile(filePath string) error {
	_, err := os.Create(filePath)
	if err != nil {
		return err
	}
	return nil
}

func copyFile(fileToCopyPath string, destinationPath string) error {
	src, err := os.Open(fileToCopyPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}
	return nil
}

func getAllFilesPathRecursively(dirPath string) ([]string, error) {
	filesNames := []string{}

	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			filesNames = append(filesNames, path)
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}

	return filesNames, nil
}

func renameFile(filePath string, newName string) error {
	err := os.Rename(filePath, newName)
	if err != nil {
		return err
	}

	return nil
}

func deleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
