package printers

import (
	"fmt"
	"os"
	"text/tabwriter"

	savem "didockerf/savesManagement"
)

const columnSeparator = "\t"

func PrintSavedDockerfilesInfos(savedDockerfileInfos []savem.SavedFileInfo) {
	printNameAndTagOfSavedFiles("DOCKERFILE", savedDockerfileInfos)
}

func PrintSavedComposeFilesInfos(savedComposeFilesInfos []savem.SavedFileInfo) {
	printNameAndTagOfSavedFiles("COMPOSE_FILE", savedComposeFilesInfos)
}

func printNameAndTagOfSavedFiles(savesType string, savedFilesInfos []savem.SavedFileInfo) {
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)

	header := savesType + "\tTAG"
	fmt.Fprintln(writer, header)

	for i := 0; i < len(savedFilesInfos); i++ {
		currentSaveInfo := savedFilesInfos[i]

		rowOfSave := currentSaveInfo.GetName() + "\t" + currentSaveInfo.GetTag()

		fmt.Fprintln(writer, rowOfSave)
	}
	writer.Flush()
}

func PrintHelpMessage(helpMsg string) {
	fmt.Println(helpMsg)
}
