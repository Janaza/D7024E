package D7024E

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Network struct {
	IP      string
	Port    int
	Contact *Contact
	Kad     *Kademlia
}

func InitNode(ip string, port int, me *Contact) *Network {
	network := &Network{
		IP:      ip,
		Port:    port,
		Contact: me,
		Kad:     InitKad(*me),
	}
	return network
}

type response struct {
	servr *net.UDPConn
	resp  *net.UDPAddr
}

const ID_INDEX = 40

//creates the content of ping
func PingMsg(contact *Contact) []byte {
	msg := []byte("ping " + contact.ID.String() + " " + contact.Address)
	return msg
}

//handles ping msgs
func HandlePingMsg(msg []byte, resp response) Contact {
	contactID := NewKademliaID(string(msg))
	ipAndPort := ipToString(msg)
	ipAndPortstring := strings.Split(ipAndPort, ":")
	ip, port := ipAndPortstring[0], ipAndPortstring[1]
	contactAdrs := ip
	contactPort := port
	contact := NewContact(contactID, contactAdrs+":"+contactPort)

	reply := []byte("PONG")
	_, err := resp.servr.WriteToUDP(reply, resp.resp)
	ErrorHandler(err)
	return contact
}

func ipToString(array []byte) string {
	hej := string(array[1+ID_INDEX:])
	return hej

}

func (network *Network) msgHandle(msg []byte, resp response) Contact {
	var returnContact Contact
	switch {
	case string(msg[:4]) == "ping":
		returnContact = HandlePingMsg(msg[5:], resp)
	case string(msg[:4]) == "find":
		fmt.Println("find node")
	default:
		returnContact = NewContact(nil, "")
	}
	return returnContact
}

func (network *Network) Listen(contact Contact, port int) {
	fmt.Println("Kademlia listener is starting...")
	listenAdrs, err := net.ResolveUDPAddr("udp", contact.Address)
	ErrorHandler(err)
	servr, err := net.ListenUDP("udp", listenAdrs)
	ErrorHandler(err)
	defer servr.Close()
	fmt.Println("Listening on: " + listenAdrs.String() + " " + contact.ID.String())
	for {
		msgbuf := make([]byte, 2048)
		n, resp, err := servr.ReadFromUDP(msgbuf)
		ErrorHandler(err)
		Response := &response{
			servr: servr,
			resp:  resp,
		}
		handledContact := network.msgHandle(msgbuf[:n], *Response)
		fmt.Println("Msg from a friend: ", handledContact)

	}
}
func (network *Network) SendPingMessage(contact *Contact) {
	me := network.Contact
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
	fmt.Println(string(respmsg[:n]))
}

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
