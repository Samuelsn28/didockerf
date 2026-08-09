package printers

import (
	"fmt"
	"os"
	"text/tabwriter"

	savem "didockerf/savesManagement"
)

const columnSeparator = "\t"

func PrintSavedDockerfilesInfos(savedDockerfileInfos []savem.SavedFileInfo) {
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)

	fmt.Fprintln(writer, "DOCKERFILE\tTAG")

	for i := 0; i < len(savedDockerfileInfos); i++ {
		currentDockerfileInfo := savedDockerfileInfos[i]

		rowOfDockerfile := currentDockerfileInfo.GetName() + "\t" + currentDockerfileInfo.GetTag()

		fmt.Fprintln(writer, rowOfDockerfile)
	}
	writer.Flush()
}
