package out

import (
	"fmt"
	"os"
)

func PanicError(err error) {
	fmt.Println(err.Error())

	panic(err.Error())
}

func PrintPanicError(msg string) {
	fmt.Println(msg)

	panic(msg)
}

func FatalError(err error) {
	fmt.Println(err.Error())

	os.Exit(1)
}

func PrintFatalError(msg string) {
	fmt.Println(msg)

	os.Exit(1)
}

func Warn(msg string) {
	fmt.Println(msg)
}
