package D7024E

import (
	"fmt"
	"log"
	"net"
	"strconv"
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

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//creates the content of ping
func PingMsg(contact *Contact) []byte {
	msg := []byte("ping " + contact.ID.String() + " " + contact.Address)
	return msg
}

func FindNodeMsg(contact *Contact) []byte {
	msg := []byte("FIND_NODE " + contact.ID.String())
	return msg
}

//handles incoming ping msgs
func (network *Network) HandlePingMsg(msg []byte, resp response) []Contact {
	fmt.Println(string(msg))
	contactID := NewKademliaID(string(msg[:ID_INDEX]))
	ipAndPort := ipToString(msg)
	ipAndPortstring := strings.Split(ipAndPort, ":")
	ip, port := ipAndPortstring[0], ipAndPortstring[1]
	contactAdrs := ip
	contactPort := port
	contact := NewContact(contactID, contactAdrs+":"+contactPort)

	reply := []byte("PONG " + network.Contact.ID.String() + " " + network.Contact.Address)
	_, err := resp.servr.WriteToUDP(reply, resp.resp)
	ErrorHandler(err)
	contactArr := make([]Contact, 1)
	contactArr[0] = contact
	return contactArr
}

func (network *Network) HandleFindNodeMsg(msg []byte, resp response) []Contact {
	closeContactsArr := network.Kad.Rtable.FindClosestContacts(NewKademliaID(string(msg[:40])), 20)

	closeCToByte := make([]byte, 0)
	closeContactsByte := make([]byte, 0)

	for i := 0; i < len(closeContactsArr); i++ {
		closeCToByte = []byte(closeContactsArr[i].ID.String() + " " + closeContactsArr[i].Address + " " + closeContactsArr[i].distance.String() + "\n")
		closeContactsByte = append(closeContactsByte, closeCToByte[:]...)
	}

	returnMsg := make([]Contact, 1)
	returnMsg[0] = NewContact(nil, " "+strconv.Itoa(len(closeContactsArr))+" contacts.")
	_, err := resp.servr.WriteToUDP(closeContactsByte, resp.resp)
	ErrorHandler(err)
	return returnMsg
}
func HandlePongMsg(msg []byte) Contact {
	contactID := NewKademliaID(string(msg[:ID_INDEX]))
	ipAndPort := ipToString(msg)
	ipAndPortstring := strings.Split(ipAndPort, ":")
	ip, port := ipAndPortstring[0], ipAndPortstring[1]
	contactAdrs := ip
	contactPort := port
	contact := NewContact(contactID, contactAdrs+":"+contactPort)
	return contact
}

func ipToString(array []byte) string {
	ipString := string(array[1+ID_INDEX:])
	return ipString

}

func (network *Network) msgHandle(msg []byte, resp response) []Contact {
	var returnContact []Contact

	switch {
	case string(msg[:4]) == "ping":
		returnContact = network.HandlePingMsg(msg[5:], resp)
		network.Kad.Rtable.AddContact(returnContact[0])
	case string(msg[:9]) == "FIND_NODE":
		returnContact = network.HandleFindNodeMsg(msg[10:], resp)
	case string(msg[:10]) == "FIND_VALUE":
		returnContact = network.HandleFindDataMsg(msg[11:], resp)
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
	fmt.Println("Listening on: " + listenAdrs.String() + " " + contact.ID.String() + "\n\n")
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
		fmt.Println("\nResponded with:  ", handledContact)
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

	if string(respmsg[:4]) == "PONG" {
		fmt.Println("Adding the following PONGER to a Bucket...")
		pongContact := HandlePongMsg(respmsg[5:n])
		network.Kad.Rtable.AddContact(pongContact)
		fmt.Println(pongContact.ID.String() + " " + pongContact.Address)
	}
}

func (network *Network) SendFindContactMessage(contact *Contact, found chan []Contact) {
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

func (network *Network) SendFindDataMessage(hash string, contact *Contact) {

	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	ErrorHandler(err)
	defer connection.Close()
	msg := FindDataMsg(contact)

	_, err = connection.Write(msg)
	ErrorHandler(err)
	fmt.Println("sent: " + string(msg))
	respmsg := make([]byte, 2048)
	n, err := connection.Read(respmsg)
	ErrorHandler(err)

	/*
	TODO:
	value := hashed data value

	if respmsg[:n] == value{
		fmt.Println("Found value: ")
		fmt.Println(string(respmsg[:n]))
		
		TODO:
		Store the value in the node
		SendStoreMessage(respmsg[:n])

	*/
	}


	else{
			fmt.Println("Got following contacts: ")
			fmt.Println(string(respmsg[:n]))
			// TODO: SendFindDataMessage to all contacts
		}

}

func FindDataMsg(contact *Contact) []byte {

	/*
	TODO:
	msg := []byte("FIND_VALUE " + The value you are looking for)
	*/
	return msg

}

func (network *Network) HandleFindDataMsg(msg []byte, resp response) []Contact {

	/*
	TODO:
	Search for value

	if value stored in node{
		return the value
	}

	else{
		closeContactsArr := network.Kad.Rtable.FindClosestContacts(NewKademliaID(string(msg[:40])), 20)

		closeCToByte := make([]byte, 0)
		closeContactsByte := make([]byte, 0)

		for i := 0; i < len(closeContactsArr); i++ {
			closeCToByte = []byte(closeContactsArr[i].ID.String() + " " + closeContactsArr[i].Address + "\n") //seperate contacts by newline
			closeContactsByte = append(closeContactsByte, closeCToByte[:]...)
		}

		returnMsg := make([]Contact, 1)
		returnMsg[0] = NewContact(nil, " "+strconv.Itoa(len(closeContactsArr))+" contacts.")
		_, err := resp.servr.WriteToUDP(closeContactsByte, resp.resp)
		ErrorHandler(err)
		return returnMsg
	}
	 */
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO STORE
}
func (network *Network) IterativeFindNode() {
	result := make(chan []Contact)
	go network.Kad.LookupContact(*network, result, *network.Contact)
	done := <-result
	for _, c := range done {
		network.Kad.Rtable.AddContact(c)
	}
	fmt.Printf("\nIterativeFindNode done, added %d contacts\n", len(done))
}

func ByteToContact(msg []byte) []Contact {
	s := string(msg)
	slice := strings.Split(s, "\n")
	arr := make([]Contact, 20)
	//var contact Contact
	for i, line := range slice {
		if len(line) != 0 {
			contact := NewContact(NewKademliaID(line[:40]), line[41:41+strings.Index(line[41:], " ")])
			if len(line[41+strings.Index(line[41:], " "):]) > 2 {
				contact.distance = NewKademliaID(line[41+strings.Index(line[41:], " "):])
			}
			arr[i] = contact
		}
	}
	return arr
}
