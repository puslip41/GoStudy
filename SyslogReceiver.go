package main

import (
	"os"
	"net"
	"fmt"
	"time"
	"runtime"
	"bufio"

	"github.com/puslip41/GoStudy/third"
)

const MINUTE_FORMAT = "200601021504"
const SECOND_FORMAT = "20060102150405"
const UDP_READ_BUFFER_SIZE = 1024*1024*10
const WRITE_BUFFER_SIZE = 1024*1024

func main() {
	port, savePath := getSyslogReceiverArgs()

	listener, err := openUdpListener(port)
	if err != nil {
		PrintError(err, "cannot open udp port")
	} else {
		defer listener.Close()

		count := 0
		beforeTime := time.Now().Format(MINUTE_FORMAT)
		buffer := make([]byte, 1024)
		var logWriter *third.LogWriter

		openLogFile := func (fileTime string) { // file close & open
			if logWriter != nil {
				logWriter.Close()
			}

			logWriter, err = openNewLogWriter(savePath, fileTime)
			if err != nil {
				PrintError(err, "cannot open log file")
			}
		}

		openLogFile(beforeTime)

		go func() { // print receive + write log count per second
			lastCheckSecond := time.Now().Format(SECOND_FORMAT)

			for {
				currentCheckSecond := time.Now()
				if lastCheckSecond != currentCheckSecond.Format(SECOND_FORMAT) {
					lastCheckSecond = currentCheckSecond.Format(SECOND_FORMAT)
					fmt.Printf("%s receive count : %d%s", currentCheckSecond.Format("2006/01/02 15:04:05"), count, getNewLineSymbol())
					count = 0
				} else {
					time.Sleep(1)
				}
			}
		}()

		for {
			length, saddr, err := listener.ReadFromUDP(buffer)
			PrintError(err, "cannot receive message")
			currentTime := time.Now()

			if beforeTime != currentTime.Format(MINUTE_FORMAT) { // change minute log file
				beforeTime = currentTime.Format(MINUTE_FORMAT)
				openLogFile(beforeTime)
			}

			logWriter.WriteFormat("%s|%s|%s%s", currentTime.Format(SECOND_FORMAT), saddr.IP.String(), buffer[:length], getNewLineSymbol() )
			count++
		}
	}

	fmt.Println(savePath)
}

func PrintError(e error, message string) {
	if e != nil {
		fmt.Printf("ERROR: %s%s%s", message, getNewLineSymbol(), e.Error())
	}
}

func getNewLineSymbol() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	} else {
		return "\n"
	}
}

func openNewLogWriter(saveDirectory, fileName string) (*third.LogWriter, error ) {
	newLogFileName := fmt.Sprintf("%s\\%s.log", saveDirectory, fileName)

	file, err := os.OpenFile(newLogFileName, os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}

	writer := bufio.NewWriterSize(file, WRITE_BUFFER_SIZE)

	return &(third.LogWriter{File: file, Writer: writer}), nil
}

func getSyslogReceiverArgs() (port string, savePath string) {
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

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	err = conn.SetReadBuffer(UDP_READ_BUFFER_SIZE)
	if err != nil {
		PrintError(err, "cannot setup udp read buffer size")
	}

	return conn, nil
}