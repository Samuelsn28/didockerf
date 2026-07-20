package arguments

func GetSubCommands() map[ArgumentId]Argument {
	export := GetSubCommandExport()
	importCmd  := GetSubCommandImport()
	dfile := GetSubCommandDfile()
	composeFile := GetSubCommandComposeFile()

	return map[ArgumentId]Argument{
		composeFile.Id: composeFile,
		dfile.Id: dfile,
		export.Id: export,
		importCmd.Id: importCmd, 
	}
}


