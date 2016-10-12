package main

import (
	"fmt"
	"os"
	"io"
)

func checkArgs(args []string) (bool, string, string, string) {
	isStart := false
	var command string
	var sourceFileName string
	var destinationFileName string

	if len(args) == 1 {
		fmt.Printf("%s: missing file operand", args[0] )
		os.Exit(-1)
	} else if len(args) == 2 {
		fmt.Printf("%s: missing destination file operand after '%s'", args[0], args[1])
		os.Exit(-1)
	} else if len(args) == 3 {
		isStart = true

		command = args[0]
		sourceFileName = args[1]
		destinationFileName = args[2]
	}

	return isStart, command, sourceFileName, destinationFileName
}

func main() {
	isStart, command, sourceFileName, destinationFileName := checkArgs(os.Args)

	if isStart {
		fi, err := os.Open(sourceFileName)
		if err != nil {
			fmt.Printf("%s: cannot stat '%s'", command, sourceFileName)
			os.Exit(-1)
		}
		defer fi.Close()

		fo, err := os.Create(destinationFileName)
		if err != nil {
			fmt.Printf("%s: cannot stat '%s'", command, destinationFileName)
			os.Exit(-1)
		}
		defer fo.Close()

		buffer := make([]byte, 1024)

		for {
			readCount, err := fi.Read(buffer)
			if err != nil && err != io.EOF {
				fmt.Printf("%s: cannot copy file: read error '%s'", command, sourceFileName)
				os.Exit(-1)
			}

			if readCount == 0 {
				break
			}

			_, err = fo.Write(buffer[:readCount])
			if err != nil {
				fmt.Printf("%s: cannot copy file: write error '%s'", command, destinationFileName)
			}
		}
	}
}
