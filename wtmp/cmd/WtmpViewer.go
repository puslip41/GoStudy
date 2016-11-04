package main

import (
	"os"
	"log"
	"io"
	"fmt"
	"errors"
	"github.com/puslip41/GoStudy/wtmp"
	"time"
)


func main() {
	command, filename, isTailing, err := getExecuteArguments()
	if err != nil {
		printCommandUsage(command)
	}

	file, err := os.OpenFile(filename, os.O_RDONLY, os.FileMode(644))
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	b := make([]byte, 384)
	utmp := wtmp.Utmp{}

	if isTailing {
		file.Seek(int64(-1 * len(b) * 5), os.SEEK_END)
	}

	for {
		if _, err := file.Read(b); err != nil {
			if err == io.EOF {
				if isTailing {
					time.Sleep(100)
				} else {
					log.Println("End File Read")
					break;
				}
			} else {
				log.Println(err)
			}
		} else {
			if err := wtmp.Unmashal(b, &utmp); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(utmp.String())
			}
		}
	}

}

func getExecuteArguments() (string, string, bool, error) {
	isTailing := false
	filename := "/var/log/wtmp"
	//filename := `D:\SourceCode\Go\src\github.com\puslip41\GoStudy\wtmp\cmd\wtmp`
	var err error
	command := os.Args[0]

	switch len(os.Args) {

	case 2:
		if os.Args[1] == "-t" {
			isTailing = true
		} else {
			filename = os.Args[1]
		}
		break

	case 3:
		if os.Args[1] == "-t" {
			isTailing = true
			filename = os.Args[2]
		} else {
			err = errors.New("Invalid Arguments")
		}
		break

	default:
		err = errors.New("Invalid Arguments")
		break
	}

	return command, filename, isTailing, err
}

func printCommandUsage(command string) {
	fmt.Printf("Usage: %s            print /var/log/wtmp\n", command)
	fmt.Printf("   or: %s [FILE]     print specified file\n", command)
	fmt.Printf("   or: %s -t [FILE]  print appended log as the file grows\n", command)
}
