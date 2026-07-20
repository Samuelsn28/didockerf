package arguments

import (
	"fmt"
)

const importId ArgumentId = "import"

func GetSubCommandImport() Argument {
	return Argument{
		Id: importId,
		Action: importFiles,
		ValidArgs: nil,
	}
}

func importFiles(args []string) bool {
	fmt.Println("Importando...")

	return true
}

