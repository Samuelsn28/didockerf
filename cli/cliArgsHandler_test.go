package cli

import (
	"fmt"
	"os"
	"testing"

	pkgArgs "didockerf/cli/arguments"
	"didockerf/out"
)

func TestArguments(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		beforeActions   func()
		confirmExpected func() (bool, string)
	}{
		{"Saving dockefile", []string{"dfile", "save", "custom_dockerfiles/debian-plus", "debian-plus:v1"}, prepareDockerfileToSave, confirmSaveDockerfile},
		{"Changing name and tag of dockerfile", []string{"dfile", "ch", "debian-plus:v1", "debian-pro:v2"}, func() {}, confirmChangeNameAndTagOfDockerfile},
		{"Changing only tag of dockerfile", []string{"dfile", "ch", "debian-pro:v2", "debian-pro:v3-ultra"}, func() {}, confirmChangeOnlyTagOfDockerfile},
		{"List all saved dockerfiles", []string{"dfile", "ls"}, func() {}, confirmTheListOfSavedDockerfiles},
		{"Getting saved dockerfile", []string{"dfile", "get", "debian-pro:v3-ultra", "."}, func() {}, confirmThatSavedDockerfileWasRetrieved},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.beforeActions()
			remainingArgs = tc.args

			Operate(pkgArgs.GetSubCommands())

			wasResultExpected, errorMessage := tc.confirmExpected()
			if !wasResultExpected {
				t.Errorf("%s", errorMessage)
			}
		})
	}
	cleanAll()
}

func prepareDockerfileToSave() {
	dockerfilesToTestDirPath := "custom_dockerfiles"
	os.Mkdir(dockerfilesToTestDirPath, 0o755)

	dockerfileContent := []byte("FROM debian:trixie")
	testDockerfilePath := dockerfilesToTestDirPath + "/" + "debian-plus"
	err := os.WriteFile(testDockerfilePath, dockerfileContent, 0o644)
	if err != nil {
		out.PrintFatalError("It was not possble to prepare save dockerfile test scenario")
	}
}

func confirmSaveDockerfile() (bool, string) {
	savedDockerfileInfo, errGetSavedDockerfileInfo := os.Stat("saves/dockerfiles/debian-plus/dockerfile_debian-plus_v1")

	if os.IsNotExist(errGetSavedDockerfileInfo) {
		return false, "Saved dockerfile saves/dockerfiles/debian-plus/dockerfile_debian-plus_v1 was not created."
	}

	testDockerfileInfo, _ := os.Stat("custom_dockerfiles/debian-plus")

	if savedDockerfileInfo.Size() != testDockerfileInfo.Size() {
		return false, fmt.Sprintf("Saved dockerfile is not equal to the original. \nOriginal size: %d. \nSaved dockerfile size: %d", testDockerfileInfo.Size(), savedDockerfileInfo.Size())
	}

	return true, ""
}

func confirmChangeNameAndTagOfDockerfile() (bool, string) {
	changedDockerfileInfo, errGetChangedDockerfileInfo := os.Stat("saves/dockerfiles/debian-pro/dockerfile_debian-pro_v2")

	if os.IsNotExist(errGetChangedDockerfileInfo) {
		return false, "Changed dockerfile saves/dockerfiles/debian-pro/dockerfile_debian-pro_v2 was not created."
	}

	testDockerfileInfo, _ := os.Stat("custom_dockerfiles/debian-plus")

	if changedDockerfileInfo.Size() != testDockerfileInfo.Size() {
		return false, fmt.Sprintf("Changed dockerfile is not equal to the old version. \nOld version: %d. \nChanged dockerfile: %d", testDockerfileInfo.Size(), changedDockerfileInfo.Size())
	}

	_, errGetOldDockerfilDirInfo := os.Stat("saves/dockerfiles/debian-plus")
	if os.IsExist(errGetOldDockerfilDirInfo) {
		return false, "Empty dir of the old dockerfile was not deleted: saves/dockerfiles/debian-plus."
	}

	return true, ""
}

func confirmChangeOnlyTagOfDockerfile() (bool, string) {
	changedDockerfileInfo, errGetChangedDockerfileInfo := os.Stat("saves/dockerfiles/debian-pro/dockerfile_debian-pro_v3-ultra")

	if os.IsNotExist(errGetChangedDockerfileInfo) {
		return false, "Changed dockerfile saves/dockerfiles/debian-pro/dockerfile_debian-pro_v3-ultra was not created."
	}

	_, errGetOldDockerfileInfo := os.Stat("saves/dockerfiles/debian-pro/dockerfile_debian-pro_v2")
	if os.IsExist(errGetOldDockerfileInfo) {
		return false, "Old dockerfile version was not deleted: saves/dockerfiles/debian-pro/dockerfile_debian-pro_v2."
	}

	testDockerfileInfo, _ := os.Stat("custom_dockerfiles/debian-plus")

	if changedDockerfileInfo.Size() != testDockerfileInfo.Size() {
		return false, fmt.Sprintf("Changed dockerfile is not equal to the old version. \nOld version: %d. \nChanged dockerfile: %d", testDockerfileInfo.Size(), changedDockerfileInfo.Size())
	}

	return true, ""
}

func confirmTheListOfSavedDockerfiles() (bool, string) {
	// Nothing yet

	// Print for test style
	fmt.Println("--------------------------------------------------")

	return true, ""
}

func confirmThatSavedDockerfileWasRetrieved() (bool, string) {
	retrievedDockerfileInfo, errOnGetDockerfile := os.Stat("./dockerfile_debian-pro_v3-ultra")
	if errOnGetDockerfile != nil {
		return false, "The wished dockerfile was not copied to the directory or don't have the name dockerfile_debian-pro_v3-ultra"
	}

	testDockerfileInfo, _ := os.Stat("custom_dockerfiles/debian-plus")

	if retrievedDockerfileInfo.Size() != testDockerfileInfo.Size() {
		return false, fmt.Sprintf("Retrieved dockerfile is not equal to the original version. \nOriginal version: %d. \nRetrieved dockerfile: %d", testDockerfileInfo.Size(), retrievedDockerfileInfo.Size())
	}

	return true, ""
}

func cleanAll() {
	errToRemoveDir := os.RemoveAll("custom_dockerfiles")
	if errToRemoveDir != nil {
		out.Warn("It was not possible remove test dockerfiles dir.")
	}

	errToRemoveDir = os.RemoveAll("saves")
	if errToRemoveDir != nil {
		out.Warn("It was not possible remove saves dir.")
	}

	errToRemoveRetrievedDockerfile := os.Remove("dockerfile_debian-pro_v3-ultra")
	if errToRemoveRetrievedDockerfile != nil {
		out.Warn("It was not possible remove retrieved dockerfile.")
	}
}
