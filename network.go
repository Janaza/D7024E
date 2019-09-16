package d7024e

import (
	"fmt"
	"log"
	"net"
	"strconv"
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
	//Handle any RPC from UDP
	msg := make([]byte, 1024)
	addr := net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}
	recv, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}

	//listen for udp msgs e.g. node sends PING, respond PONG
	for {
		log.Printf("listening...")
		n, returnAddr, err := recv.ReadFromUDP(msg)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("received %v bytes, ret addr %v, msg %s", n, returnAddr, string(msg[:n]))

		//check if pinged if msg == ping etc...
		if string(msg[:4]) == "PING" {
			reply := []byte(fmt.Sprintf("PONG"))
			n, err = recv.WriteToUDP(reply, returnAddr)
			if err != nil {
				log.Fatal(err)
			}
		}

		//if string(msg[:n]) == "SendFindContactMessage" + "contact" call func SendFindContactMessage etc...

	}

}

//Just an example of sending a message to func Listen
func (network *Network) SendPing(contact *Contact) {
	host, port, err := net.SplitHostPort(contact.Address)
	if err != nil {
		log.Fatal(err)
	}
	iport, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(err)
	}
	addr := net.UDPAddr{
		IP:   net.ParseIP(host),
		Port: iport,
	}

	conn, err := net.DialUDP("udp", nil, &addr)
	if err != nil {
		log.Fatal(err)
	}
	n, err := conn.Write([]byte("PING " + (network.me.ID.String()))) //send ping with my KadID
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sent: "+string([]byte("PING "))+network.me.ID.String()+" %d bytes \nTo ip "+contact.Address+":"+"%d", n, addr.Port)
	msg := make([]byte, 1024)
	n, err = conn.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("server replied with: %s \n", string(msg[:n]))
	return
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
