package file

import (
	"fmt"
	"io"
	"os"
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
		fmt.Println("Erro ao create file")
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

