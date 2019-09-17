package d7024e

import (
	"fmt"
	"net"
	"strings"
)


type Network struct {
	IP string
	Port int
	Contact *Contact
}

const ID_INDEX  = 40

//creates the content of ping
func PingMsg(contact *Contact) []byte{
	msg := []byte(contact.ID.String())
	msg = append(msg, contact.Address...)
	return msg
}
//handles ping msgs
func HandlePingMsg(msg []byte) Contact{
	contactID := NewKademliaID(string(msg[0:ID_INDEX]))
	ipAndPort := ipToString(msg)
	ipAndPortstring := strings.Split(ipAndPort, ":")
	ip, port := ipAndPortstring[0], ipAndPortstring[1]
	contactAdrs := ip
	contactPort := port
	contact := NewContact(contactID, contactAdrs +":"+ contactPort)
	fmt.Println(contact.ID.String())
	fmt.Println(contact.Address)
	return contact
}

func ipToString(array []byte) string{
	hej := string(array[ID_INDEX:])
	return hej
}


func msgHandle(arg arg){
	switch:
		case "ping"{
			do thing
	}
		case "find"{

	}
}

func Listen(contact Contact, port int) {
	fmt.Println("Kademlia listener is starting...")
	listenAdrs, err := net.ResolveUDPAddr("udp", contact.Address)
	ErrorHandler(err)
	servr, err := net.ListenUDP("udp", listenAdrs)
	ErrorHandler(err)
	defer servr.Close()
	for {
		fmt.Println("Listening on: " + listenAdrs.String() +" "+ contact.ID.String())
		msgbuf := make([]byte, 2048)
		_ ,_ , err := servr.ReadFrom(msgbuf)
		ErrorHandler(err)
		handledContact := HandlePingMsg(msgbuf)
		fmt.Println("Msg from a friend: ", handledContact)

	}
}
func (network *Network) SendPingMessage(me *Contact, contact *Contact) {
	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	ErrorHandler(err)
	defer connection.Close()
	msg := PingMsg(me)
	_, err = connection.Write(msg)
	ErrorHandler(err)

}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
