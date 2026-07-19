package arguments

type ArgumentId string

type Argument struct {
	Id ArgumentId
	Action func(remainingCliArgs []string) (stopReadingNextArgs bool)
	ValidArgs map[ArgumentId]Argument
}


