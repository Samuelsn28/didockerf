package file

import (
	"os"
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
