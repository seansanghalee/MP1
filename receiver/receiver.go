package receiver

import (
	"MP1/packet"
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func Receiver(x string) {
	for {
		l, err := net.Listen("tcp", ":"+x)
		if err != nil {
			fmt.Println(err)
			return
		}

		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		dec := gob.NewDecoder(c) // read from the channnel
		var p packet.Packet
		dec.Decode(&p)
		Unicast_receive(c.RemoteAddr(), p.Message)
		c.Close()
		l.Close()
	}
}

// Unicast_receive delivers the message received from the source process
func Unicast_receive(source net.Addr, message string) {

	t := time.Now()
	myTime := t.Format(time.RFC3339) + "\n"
	fmt.Printf("Received [%s] from process %s, system time is %s", message, source, myTime) // if we can, fix it so that process id shows up

	return
}