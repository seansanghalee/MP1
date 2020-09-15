package main

import (
	"MP1/config"
	"MP1/receiver"
	"MP1/sender"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		ID, IP, PORT string
	)

	arguments := os.Args // takes user input for ID, IP and Port number for current process
	if len(arguments) == 1 {
		fmt.Println("Please provide ID:IP:Port")
		return
	}
	temp := strings.Split(arguments[1], ":")
	ID, IP, PORT = temp[0], temp[1], temp[2]
	fmt.Printf("Process ID: %v, IP: %v, PORT: %v\n", ID, IP, PORT)

	var (
		minDelay, maxDelay int
		ids, ips, ports    []string
	)

	file := config.ReadConfig()
	minDelay, maxDelay, ids, ips, ports = config.Extract(file)
	config.Display(minDelay, maxDelay, ids, ips, ports) //Displays config for UI benefit

	go receiver.Receiver(PORT) // thread 1 : listen and receive

	for {
		sender.Sender(ID) // thread 2 : dial and send
	}

}
