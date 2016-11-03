package main

import (
	"os"
	"log"
	"io"
	"fmt"
	"errors"
	"github.com/puslip41/GoStudy/wtmp"
)


func main() {
	filename, isTailing, err := getExecuteArguments()

	file, err := os.OpenFile(filename, os.O_RDONLY, os.FileMode(644))
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	b := make([]byte, 384)
	utmp := wtmp.Utmp{}

	if isTailing {

	} else {
		for {
			if _, err := file.Read(b); err != nil {
				if err == io.EOF {
					log.Println("End File Read")
					break;
				}
				log.Println(err)
			} else {
				if err := wtmp.Unmashal(b, &utmp); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(utmp.String())
				}
			}
		}
	}



}

func getExecuteArguments() (string, bool, error) {
	isTailing := false
	//filename := "/var/log/wtmp"
	filename := `D:\SourceCode\Go\src\github.com\puslip41\GoStudy\wtmp\cmd\wtmp`
	var err error

	switch len(os.Args) {

	case 2:
		filename = os.Args[1]
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
	}

	return filename, isTailing, err
}
