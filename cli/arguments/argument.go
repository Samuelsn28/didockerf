package arguments

type ArgumentID string

type Argument struct {
	ID        ArgumentID
	Action    func(remainingCliArgs []string) (stopReadingNextArgs bool)
	ValidArgs map[ArgumentID]Argument
}
