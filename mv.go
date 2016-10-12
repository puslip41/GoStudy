package main

import (
	"fmt"
	"os"
	"syscall"
)

func checkArgs(args []string) (bool, string, string, string) {
	isValid := false

	var command, sourceFileName, destinationFileName string

	argsLen := len(args)

	if argsLen == 1 {
		fmt.Println("%s: missing file operand", args[0] )
	} else if argsLen == 2 {
		fmt.Println("%s: missing destination file operand after '%s'", args[0], args[1])
	} else if argsLen == 3 {
		isValid = true

		command = args[0]
		sourceFileName = args[1]
		destinationFileName = args[2]
	}

	return isValid, command, sourceFileName, destinationFileName
}

func main() {
	isValidArgs, command, sourceFileName, destinationFileName := checkArgs(os.Args)

	if isValidArgs {
		err := os.Link(sourceFileName, destinationFileName)
		if err != nil {
			fmt.Println("%s: cannot stat '%s'", command, sourceFileName)
			os.Exit(-1)
		}

		err = syscall.Unlink(sourceFileName)
		if err != nil {
			fmt.Println("%s: %s", command, err.Error())
			os.Exit(-1)
		}
	}
}
