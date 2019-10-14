package D7024E

import (
	"fmt"
	"net"
)

//RPC senders

func (kademlia *Kademlia) SendPingMessage(contact *Contact) {
	me := kademlia.Net.Contact
	//me := network.Contact
	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	ErrorHandler(err)
	defer connection.Close()
	msg := PingMsg(me)
	_, err = connection.Write(msg)
	ErrorHandler(err)
	fmt.Println("sent: " + string(msg))
	respmsg := make([]byte, 1024)
	n, err := connection.Read(respmsg)
	ErrorHandler(err)

	if string(respmsg[:4]) == "PONG" {
		fmt.Println("Adding the following PONGER to a Bucket...")
		pongContact := HandlePongMsg(respmsg[5:n])
		kademlia.Rtable.AddContact(pongContact)
		//network.Kad.Rtable.AddContact(pongContact)
		fmt.Println(pongContact.ID.String() + " " + pongContact.Address)
	}
}

func (kademlia *Kademlia) SendFindContactMessage(contact *Contact, found chan []Contact) {
	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	ErrorHandler(err)
	defer connection.Close()
	msg := FindNodeMsg(contact)

	_, err = connection.Write(msg)
	ErrorHandler(err)
	fmt.Println("sent: " + string(msg) + " to: " + contact.Address)
	respmsg := make([]byte, 2048)
	n, err := connection.Read(respmsg)
	ErrorHandler(err)
	c := ByteToContact(respmsg[:n])
	found <- c

}

func (kademlia *Kademlia) SendStoreMessage(contact *Contact, data []byte) {
	//me := network.Contact
	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	ErrorHandler(err)
	defer connection.Close()
	msg := StoreMsg(data)
	_, err = connection.Write(msg)
	ErrorHandler(err)
	respmsg := make([]byte, 1024)
	n, err := connection.Read(respmsg)
	ErrorHandler(err)
	fmt.Println(string(respmsg[:n]))
}


func (kademlia *Kademlia) SendFindDataMessage(hash string, contact *Contact, found chan []Contact, value chan string) {
	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	ErrorHandler(err)
	defer connection.Close()
	msg := FindDataMsg(hash)
	_, err = connection.Write(msg)
	ErrorHandler(err)


	//fmt.Println("sent: " + string(msg) + " to: " + contact.Address)
	respmsg := make([]byte, 2048)
	n, err := connection.Read(respmsg)
	ErrorHandler(err)


	c := make([]Contact, 0)
	if string(respmsg[:2]) == "OK"{
		found <- c
		value <- string(respmsg[:n])
	} else{
		c = ByteToContact(respmsg[:n])
		found <- c
		value <- ""
	}
}