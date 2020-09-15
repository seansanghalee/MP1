package packet

// Packet contains the ID of the sender and the message to send
type Packet struct {
	Source, Message string
}

// Construct allows user to create a struct type Packet
func Construct(s, m string) Packet {
	packet := Packet{s, m}
	return packet
}
