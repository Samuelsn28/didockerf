package arguments

import (
	"fmt"
)

func GetSubCommandImport() Argument {
	return Argument{
		Id: "import",
		Action: importFiles,
		ValidArgs: nil,
	}
}

func importFiles(args []string) bool {
	fmt.Println("Importando...")

	return true
}

