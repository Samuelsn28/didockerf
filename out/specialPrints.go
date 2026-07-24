package out

import (
	"fmt"
	"os"
	"text/tabwriter"

	"didockerf/model"
)

const columnSeparator = "\t"

func PrintDockerfilesInTable(dockerfiles []model.Dockerfile) {
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)

	fmt.Fprintln(writer, "DOCKERFILE\tTAG")

	for i := 0; i < len(dockerfiles); i++ {
		currentDockerfile := dockerfiles[i]

		rowOfDockerfile := currentDockerfile.Name + "\t" + currentDockerfile.Version

		fmt.Fprintln(writer, rowOfDockerfile)
	}
	writer.Flush()
}

