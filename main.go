package main

import (
	"didockerf/cli"
	pkgArgs "didockerf/cli/arguments"
	"didockerf/model"
	"didockerf/out"
)

func main() {
	
	tabela := model.CreateTable(3, 3)
	tabela.SetValue(1, 1, "Esquerda autoritária")
	tabela.SetValue(1, 2, "Autoritário")
	tabela.SetValue(1, 3, "Direita autoritária")
	tabela.SetValue(2, 1, "Esquerda")
	tabela.SetValue(2, 2, "Centrão")
	tabela.SetValue(2, 3, "Direita")
	tabela.SetValue(3, 1, "Esquerda liberal")
	tabela.SetValue(3, 2, "Libertário")
	tabela.SetValue(3, 3, "Direita liberal")

	out.PrintTable(tabela)

	determineSubCommand()
}

func determineSubCommand() {
	subCommands := pkgArgs.GetSubCommands()

	cli.Operate(subCommands)


}



