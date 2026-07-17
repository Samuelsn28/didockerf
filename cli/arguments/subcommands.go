package arguments

func GetSubCommands() map[ArgumentId]Argument {
	export := GetSubCommandExport()
	importCmd  := GetSubCommandImport()
	dfile := GetSubCommandDfile()

	return map[ArgumentId]Argument{
		dfile.Id: dfile,
		export.Id: export,
		importCmd.Id: importCmd, 
	}
}


