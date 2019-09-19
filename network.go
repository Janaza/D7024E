package D7024E

import (
	"fmt"
	"log"
	"net"
	"strings"
	"strconv"
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

func FindNodeMsg(contact *Contact) []byte {
 	msg := []byte("FIND_NODE " + contact.ID.String())
 	return msg
 }

//handles ping msgs
func HandlePingMsg(msg []byte, resp response) []Contact {
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
	contactArr := make([]Contact, 1)

	return append(contactArr, contact)
}

func (network *Network) HandleFindNodeMsg(msg []byte, resp response) []Contact {
	/*
	contactID := NewKademliaID(string(msg))
	ipAndPort := ipToString(msg)
	ipAndPortstring := strings.Split(ipAndPort, ":")
	ip, port := ipAndPortstring[0], ipAndPortstring[1]
	contactAdrs := ip
	contactPort := port
	contact := NewContact(contactID, contactAdrs+":"+contactPort)
	*/



    closeContactsArr := network.Kad.Rtable.FindClosestContacts(NewKademliaID(string(msg[:40])), 20)

    closeCToByte := make([]byte, 20)
    closeContactsByte := make([]byte, len(closeContactsArr))

    var closeC Contact
    closeContacts := make([]Contact, 20)
    Arrlength := strconv.Itoa(len(closeContactsArr))
    fmt.Println("TESTMEDDDDDDDDDDDDDDDDDDDDDDDEEEEEEEEEEEE" + Arrlength)
    for i := 0; i < len(closeContactsArr); i++{

        closeCToByte = []byte(closeContactsArr[i].ID.String() + " " + closeContactsArr[i].Address)
        closeContactsByte = append(closeContactsByte, closeCToByte[:]...)
        fmt.Println(string(closeContactsByte))
        closeC = NewContact(closeContactsArr[i].ID, closeContactsArr[i].Address)
        closeContacts = append(closeContacts, closeC)
    }


	reply := closeContacts
	_, err := resp.servr.WriteToUDP(closeContactsByte, resp.resp)
	ErrorHandler(err)
	return reply
}

func ipToString(array []byte) string {
	hej := string(array[1+ID_INDEX:])
	return hej

}

func (network *Network) msgHandle(msg []byte, resp response) []Contact {
	var returnContact []Contact

	switch {
	case string(msg[:4]) == "ping":
		returnContact = HandlePingMsg(msg[5:], resp)
	case string(msg[:9]) == "FIND_NODE":
	    fmt.Println(string(msg[10:]))
		returnContact = network.HandleFindNodeMsg(msg[10:], resp)
	default:
		returnContact = append(returnContact, NewContact(nil, ""))
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
		fmt.Println("Msg from a friend: ", string(msgbuf[:n]))
		fmt.Println("Responded with:  ", handledContact)

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
	//me := network.Contact
   	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
   	connection, err := net.DialUDP("udp", nil, RemoteAddress)
   	ErrorHandler(err)
   	defer connection.Close()
    msg := FindNodeMsg(contact)



   	_, err = connection.Write(msg)
   	ErrorHandler(err)
   	fmt.Println("sent: " + string(msg))
   	respmsg := make([]byte, 1024)
   	_, err = connection.Read(respmsg)
   	ErrorHandler(err)
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO FIND_VALUE
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO STORE
}
