package main

import (
	"fmt"
	"os"
	"time"
)

func getTailArgs() (bool, string) {
	isValid := false
	argsLen := len(os.Args)

	var fileName string

	if argsLen == 1 {
		fmt.Println("%s: missing file operand", os.Args[0] )
	} else if argsLen == 2 {
		isValid = true

		fileName = os.Args[1]
	}

	return isValid, fileName
}

func readLine(f *os.File)  string {
	var buffer []byte
	readBuffer := make([]byte, 1)

	for true {
		i, err := f.Read(readBuffer)
		if err != nil {
			time.Sleep(1)
		} else {
			if i == 1 {
				if readBuffer[0] == '\n' {
					break
				} else {
					buffer = append(buffer, readBuffer[0])
				}
			}
		}
	}

	return string(buffer)
}

func main() {
	isValidArgs, readFileName := getTailArgs()

	if isValidArgs {
		f, err := os.Open(readFileName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		defer f.Close()

		f.Seek(0, os.SEEK_END)

		for true {
			fmt.Println(readLine(f))
		}
	}
}
