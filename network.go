package d7024e

import (
	"fmt"
	"net"
	"strings"
	"log"
	"strconv"
)


type Network struct {
	IP string
	Port int
	Contact *Contact
	Kad     *Kademlia
}

func InitNode(ip string, port int, me Contact) *Network {
	network := &Network{
		IP: ip,
		Port:    port,
		Contact:      me,
		Kad:     InitKad(me),
	}
	return network
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


)



/*func Listen(ip string, port int) {
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
*/
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
