package file

import (
	"os"
	"strings"
)

func dirExist(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func createDir(dir string) error {
	err := os.MkdirAll(dir, 0755)

	return err
}

func getNameWithoutPath(fileWithPath string) string {
	fileWithPathSplitted := strings.Split(fileWithPath, "/")

	return fileWithPathSplitted[ len(fileWithPathSplitted) - 1 ]
}
