package file

import (
	"io"
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

func getPathWithoutName(fileWithPath string) string {
	fileWithPathSplitted := strings.Split(fileWithPath, "/")
	fileWithPathSplitted = fileWithPathSplitted[: (len(fileWithPathSplitted) - 1)]

	return strings.Join(fileWithPathSplitted, "/")
}

func isDirEmpty(dirPath string) bool {
	dir, _ := os.Open(dirPath)
	defer dir.Close()

	_, errReadDir := dir.Readdirnames(1)
    if errReadDir == io.EOF {
        return true
    }
    return false
}

func deleteDir(dirPath string) error {
	errDirIsntEmpty := os.Remove(dirPath)

	if (errDirIsntEmpty != nil){
		return errDirIsntEmpty
	}
	return nil
 }
