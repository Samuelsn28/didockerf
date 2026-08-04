package arguments

import (
	"fmt"
)

const importID ArgumentID = "import"

func GetSubCommandImport() Argument {
	return Argument{
		ID:        importID,
		Action:    importFiles,
		ValidArgs: nil,
	}
}

func importFiles(args []string) bool {
	fmt.Println("Importando...")

	return true
}
