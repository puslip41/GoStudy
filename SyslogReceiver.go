package main

import (
	"os"
	"net"
	"fmt"
	"time"
	"runtime"
)

const MINUTE_FORMAT = "200601021504"
const SECOND_FORMAT = "20060102150405"

func main() {
	port, savePath := getArgs()

	newlineSymbol := getNewLineSymbol()

	listener, err := openUdpListener(port)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer listener.Close()

		count := 0
		beforeTime := time.Now().Format(MINUTE_FORMAT)

		buffer := make([]byte, 1024)

		logFile, err := openNewLogFile( savePath, beforeTime )
		if err != nil {
			fmt.Println(err.Error())
		}

		go func() { // print receive + write log count per second
			lastCheckSecond := time.Now().Format(SECOND_FORMAT)

			for {
				currentCheckSecond := time.Now()
				if lastCheckSecond != currentCheckSecond.Format(SECOND_FORMAT) {
					lastCheckSecond = currentCheckSecond.Format(SECOND_FORMAT)
					fmt.Printf("%s receive count : %d%s", currentCheckSecond.Format("2006/01/02 15:04:05"), count, newlineSymbol)
					count = 0
				} else {
					time.Sleep(1)
				}
			}
		}()

		for {
			length, saddr, err := listener.ReadFromUDP(buffer)
			PrintError(err, "cannot receive message", newlineSymbol)
			currentTime := time.Now()

			if beforeTime != currentTime.Format(MINUTE_FORMAT) { // change minute log file
				beforeTime = currentTime.Format(MINUTE_FORMAT)
				logFile.Close()
				logFile, err = openNewLogFile(savePath, beforeTime)
				PrintError(err, "cannot create log file", newlineSymbol)
			}

			if logFile != nil {
				logFile.WriteString(fmt.Sprintf("%s|%s|%s%s", currentTime.Format(SECOND_FORMAT), saddr.IP.String(), buffer[:length], newlineSymbol))
				count++
			}
		}
	}

	fmt.Println(savePath)
}


func PrintError(e error, message, newline string) {
	if e != nil {
		fmt.Printf("ERROR: %s%s%s", message, newline, e.Error())
	}
}

func getNewLineSymbol() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	} else {
		return "\n"
	}
}

func openNewLogFile(saveDirectory, fileName string) (* os.File, error) {
	newLogFileName := fmt.Sprintf("%s\\%s.log", saveDirectory, fileName)
	fmt.Println("Open Log File: ", newLogFileName)
	return os.OpenFile(newLogFileName, os.O_APPEND|os.O_CREATE, 0660)
}

func getArgs() (port string, savePath string) {
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "514"
	}

	if len(os.Args) > 2 {
		savePath = os.Args[2]
	} else {
		savePath, _ = os.Getwd()
	}

	return
}

func openUdpListener(port string) (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		return nil, err
	}

	return net.ListenUDP("udp", addr)
}
