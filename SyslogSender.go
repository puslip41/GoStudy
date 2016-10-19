package main

import (
	"os"
	"fmt"
	"net"
	"time"
)

const SECOND_FORMAT = "20060102150405"

func main() {
	if len(os.Args) != 3 && len(os.Args) != 4 {
		fmt.Printf("Usage: %s [IP] [PORT]\n", os.Args[0])
		fmt.Printf("Usage: %s [IP] [PORT] [MESSAGE]\n", os.Args[0])
		os.Exit(0)
	}

	serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]))
	if err != nil {
		fmt.Printf("Error: Server Initialize Fail\n%s\n", err.Error())
		os.Exit(-1)
	}

	clientAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		fmt.Printf("Error: Client Initialize Fail\n%s\n", err.Error())
		os.Exit(-1)
	}

	conn, err := net.DialUDP("udp", clientAddr, serverAddr)
	if err != nil {
		fmt.Printf("Error: Create Socket Fail\n%s\n", err.Error())
		os.Exit(-1)
	}
	defer conn.Close()

	sendCount := 0

	message := "<14>1 2015-01-21T07:13:46.821632Z [fw4_deny] [58.29.31.35] 2015-01-21 16:13:46,2015-01-21 16:13:46,0,FW_Bmaster,0,Undefined,50.63.93.1,80,125.240.146.101,23258,TCP,EXT,1,64, ,AS,Deny by Deny Rule"
	if len(os.Args) == 4 {
		message = os.Args[3]
	}

	go func() {
		lastCheckSecond := time.Now().Format(SECOND_FORMAT)

		for {
			currentCheckSecond := time.Now()
			if lastCheckSecond != currentCheckSecond.Format(SECOND_FORMAT) {
				lastCheckSecond = currentCheckSecond.Format(SECOND_FORMAT)
				fmt.Printf("%s send count : %d\n", currentCheckSecond.Format("2006/01/02 15:04:05"), sendCount)
				sendCount = 0
			} else {
				time.Sleep(1)
			}
		}
	}()

	for i := 1; i > 0; i++ {
		conn.Write([]byte(fmt.Sprintf("%8d %s", i, message)))
		sendCount++
	}

}
