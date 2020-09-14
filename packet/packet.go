package packet

type Packet struct {
	Message string
}

func Construct(m string) Packet {
	packet := Packet{m}
	return packet
}