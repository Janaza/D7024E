package d7024e

import (
	"fmt"
	"log"
	"net"
)

//Very unsure about this struct =)
type Network struct {
	address string
	port    int
	me      Contact
}

func InitNetwork(ip string, port int, me Contact) *Network {
	network := &Network{}
	network.address = ip
	network.port = port
	network.me = me
	return network
}

func Listen(ip string, port int) {
	// TODO
	msg := make([]byte, 1024)
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ip),
	}
	recv, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}
	n, returnAddr, err := recv.ReadFromUDP(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("received %v bytes, ret addr %v, msg %s", n, returnAddr, string(msg[:n])) //check if pinged if msg = ping etc...

	//listen for udp msgs e.g. node sends PING, respond PONG?

	reply := []byte(fmt.Sprintf("PONG"))
	n, err = recv.WriteToUDP(reply, returnAddr)

}

//*Network is sender, *Contact is reciver? See main.go:40
func (network *Network) SendPingMessage(contact *Contact) {
	// TODO PING
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO FIND_NODE
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO FIND_VALUE
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO STORE
}
