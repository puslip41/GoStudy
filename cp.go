package main

import (
	"fmt"
	"os"
	"io"
)

func getCpArgs(args []string) (bool, string, string, string) {
	isValid := false
	var command, sourceFileName, destinationFileName string

	if len(args) == 1 {
		fmt.Println("%s: missing file operand", args[0] )
		os.Exit(-1)
	} else if len(args) == 2 {
		fmt.Println("%s: missing destination file operand after '%s'", args[0], args[1])
		os.Exit(-1)
	} else if len(args) == 3 {
		isValid = true

		command = args[0]
		sourceFileName = args[1]
		destinationFileName = args[2]
	}

	return isValid, command, sourceFileName, destinationFileName
}

func main() {
	isValidArgs, command, sourceFileName, destinationFileName := getCpArgs(os.Args)

	if isValidArgs {
		fi, err := os.Open(sourceFileName)
		if err != nil {
			fmt.Println("%s: cannot stat '%s'\n%s", command, sourceFileName, err.Error())
			os.Exit(-1)
		}
		defer fi.Close()

		fo, err := os.Create(destinationFileName)
		if err != nil {
			fmt.Println("%s: cannot stat '%s'", command, destinationFileName)
			os.Exit(-1)
		}
		defer fo.Close()

		buffer := make([]byte, 1024)

		for {
			readCount, err := fi.Read(buffer)
			if err != nil && err != io.EOF {
				fmt.Println("%s: cannot copy file: read error '%s'", command, sourceFileName)
				os.Exit(-1)
			}

			if readCount == 0 {
				break
			}

			_, err = fo.Write(buffer[:readCount])
			if err != nil {
				fmt.Println("%s: cannot copy file: write error '%s'", command, destinationFileName)
			}
		}
	}
}
