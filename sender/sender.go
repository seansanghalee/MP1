package sender

import (
	"MP1/config"
	"MP1/packet"
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func Sender() {
	d, c := UserInput()
	Unicast_send(d, c)
}

// UserInput
func UserInput() (string, string) {
	reader := bufio.NewReader(os.Stdin) // input = send 2 Hello
	text, _ := reader.ReadString('\n')  // text = "send 2 Hello\n"
	str := strings.Split(text, " ")     // str = ["send", "2", "Hello"]
	d := str[1]
	content := strings.TrimSpace(str[2])
	return d, content
}

// Unicast_send sends message to the destination process
func Unicast_send(destination, message string) {
	file := config.ReadConfig()
	_, _, ids, ips, ports := config.Extract(file)
	location := 0
	for i := range ids {
		if ids[i] == destination {
			location = i
		}
	}

	c := Dialer(ips[location], ports[location])
	enc := gob.NewEncoder(c)
	packet := packet.Construct(message)
	enc.Encode(packet)
	t := time.Now()
	myTime := t.Format(time.RFC3339) + "\n"
	fmt.Printf("Sent [%s] to process %s, system time is %s", message, destination, myTime)
	c.Close()

	// rand.Seed(time.Now().UnixNano())
	// var min, max int
	// duration := min + rand.Intn(max - min)
	// time.Sleep(duration * time.Millisecond)

	return
}

func Dialer(ip, port string) (conn net.Conn) {
	address := ip + ":" + port
	c, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	return c
}