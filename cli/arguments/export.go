package arguments

import (
	"fmt"
)

const exportID ArgumentID = "export"

func GetSubCommandExport() Argument {
	return Argument{
		ID:        exportID,
		Action:    exportAllFiles,
		ValidArgs: make(map[ArgumentID]Argument),
	}
}

func exportAllFiles([]string) bool {
	fmt.Println("Exportando...")

	return true
}
