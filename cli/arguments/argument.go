package arguments

type ArgumentId string

type Argument struct {
	Id ArgumentId
	Action func([]string) bool
	ValidArgs map[ArgumentId]Argument
}


