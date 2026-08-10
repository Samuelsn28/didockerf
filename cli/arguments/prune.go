package arguments

import (
	"errors"

	"didockerf/out"
	savem "didockerf/savesManagement"
)

const pruneID ArgumentID = "prune"

func GetArgPruneDockerfile() Argument {
	return makeArgPrune(pruneDockerfile)
}

func GetArgPruneComposeFile() Argument {
	return makeArgPrune(pruneComposeFile)
}

func makeArgPrune(action func([]string) bool) Argument {
	return Argument{
		ID:        pruneID,
		Action:    action,
		ValidArgs: nil,
	}
}

func pruneDockerfile(args []string) bool {
	errOnArgs := checkIfPruneDockerfilesArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
		return true
	}

	errOnPrune := savem.RemoveAllSavedDockerfiles()
	if errOnPrune != nil {
		out.FatalError(errOnPrune)
	}

	return true
}

func pruneComposeFile(args []string) bool {
	errOnArgs := checkIfPruneComposeFilesArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
		return true
	}

	errOnPrune := savem.RemoveAllSavedComposeFiles()
	if errOnPrune != nil {
		out.FatalError(errOnPrune)
	}

	return true
}

func checkIfPruneDockerfilesArgsAreCorrect(args []string) error {
	return checkIfPruneArgsAreCorrect(args)
}

func checkIfPruneComposeFilesArgsAreCorrect(args []string) error {
	return checkIfPruneArgsAreCorrect(args)
}

func checkIfPruneArgsAreCorrect(args []string) error {
	if len(args) != 0 {
		return errors.New("Prune command don't need any arguments.")
	}

	return nil
}
