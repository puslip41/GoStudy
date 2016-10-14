package main

import (
	"os"
	"strings"
)

func main(){
	isValidArgs, keyword, fileName := getGrepArgs()

	if isValidArgs {
		f, err := getInputFile(fileName)
		if err != nil {
			println(err.Error())
		} else {
			defer f.Close()

			for true {
				readLine, err := readLineGrep(f)
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

func readLineGrep(f *os.File)  (string, error) {
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
	if fileName == "" {
		return os.Stdin, nil
	} else {
		return os.Open(fileName)
	}
}

func getGrepArgs() (bool, string, string) {
	if len(os.Args) == 1 {
		println("%s: missing file operand", os.Args[0] )
		return false, "", ""
	} else if len(os.Args) == 2 {``
		return true, os.Args[1], ""
	} else if len(os.Args) == 3 {
		return true, os.Args[1], os.Args[2]
	} else  {
		return false, "", ""
	}
}

func printGrepHelp(command string) {
	println("Usage: %s [KEYWORD] [FILENAME]", command)
	println("Usage: %s [KEYWORD]", command)
}
