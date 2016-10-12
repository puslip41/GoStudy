package main

import (
	"os"
	"fmt"
	"strings"
)

func readLineTail(f *os.File)  (string, error) {
	var buffer []byte
	readBuffer := make([]byte, 1)

	for true {
		i, err := f.Read(readBuffer)
		if err != nil {
			return "", err
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

	return string(buffer), nil
}

func getInputFile(fileName string) (*os.File, error) {
	var f *os.File
	var err error
	if fileName == "" {
		f = os.Stdin
	} else {
		f, err = os.Open(fileName)
	}

	return f, err
}

func getGrepArgs() (bool, string, string) {
	isValid := false
	argsLen := len(os.Args)

	var keyword, fileName string

	if argsLen == 1 {
		fmt.Println("%s: missing file operand", os.Args[0] )
	} else if argsLen == 2 {
		isValid = true

		keyword = os.Args[1]
		fileName = ""
	} else if argsLen == 3 {
		isValid = true

		keyword = os.Args[1]
		fileName = os.Args[2]
	}

	return isValid, keyword, fileName
}

func main(){
	isValidArgs, keyword, fileName := getGrepArgs()

	if isValidArgs {
		f, err := getInputFile(fileName)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			defer f.Close()
			
			var readLine string
			for true {
				readLine, err = readLineTail(f)
				if err != nil {
					break;
				} else {
					if strings.Contains(readLine, keyword) {
						println(readLine)
					}
				}
			}
		}
	}
}
