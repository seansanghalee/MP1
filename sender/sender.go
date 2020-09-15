package sender

import (
	"MP1/config"
	"MP1/packet"
	"bufio"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

// Sender contains three methods that
// 1. asks for user input
// 2. creates a packet to send
// 3. sends over the packet through TCP channel
func Sender(ID string) {
	d, c := UserInput()
	p := packet.Construct(ID, c)
	go Unicast_send(d, p)
}

// UserInput prompts user for the ID of destination process and the message
func UserInput() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	str := strings.Split(text, " ")
	d := str[1]
	content := strings.TrimSpace(str[2])
	t := time.Now()
	myTime := t.Format(time.StampMilli) + "\n"
	fmt.Printf("Sent [%s] to process %s, system time is %s", content, d, myTime)
	return d, content
}

// Unicast_send sends message to the destination process
func Unicast_send(destination string, packet packet.Packet) {
	file := config.ReadConfig()
	min, max, ids, ips, ports := config.Extract(file)
	location := 0
	for i := range ids {
		if ids[i] == destination {
			location = i
		}
	} // gets the index of the destination process from the configuration file

	rand.Seed(time.Now().UnixNano())
	duration := min + rand.Intn(max-min)
	time.Sleep(time.Duration(duration) * time.Millisecond) // creates a delay that simulates real-life network delay

	c := Dialer(ips[location], ports[location])
	enc := gob.NewEncoder(c)
	enc.Encode(packet)
	c.Close()

	return
}

// Dialer establishes a connection with the destination tcp server
func Dialer(ip, port string) (conn net.Conn) {
	address := ip + ":" + port
	c, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	return c
}
