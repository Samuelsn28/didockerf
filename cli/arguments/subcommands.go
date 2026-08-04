package arguments

func GetSubCommands() map[ArgumentID]Argument {
	export := GetSubCommandExport()
	importCmd := GetSubCommandImport()
	dfile := GetSubCommandDfile()
	composeFile := GetSubCommandComposeFile()

	return map[ArgumentID]Argument{
		composeFile.ID: composeFile,
		dfile.ID:       dfile,
		export.ID:      export,
		importCmd.ID:   importCmd,
	}
}
