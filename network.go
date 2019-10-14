package D7024E

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	//"net/http"
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

type data struct {
	rpc  string `json:string`
	data string `json:string`
}

const ID_INDEX = 40

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//creates the content of ping
func PingMsg(contact *Contact) *data {
	//msg := []byte("ping " + contact.ID.String() + " " + contact.Address)
	msg := &data{
		rpc:  "ping",
		data: contact.ID.String() + " " + contact.Address,
	}
	return msg
}

func FindNodeMsg(contact *Contact) []byte {
	msg := []byte("FIND_NODE " + contact.ID.String())
	return msg
}

func StoreMsg(data []byte) []byte {
	msg := []byte("store " + string(data))
	return msg
}

func (network *Network) HandleStoreMsg(msg []byte, resp response) []Contact {
	hashedData := HashData(msg[0:])
	fmt.Println(hashedData)

	if _, ok := network.Kad.hashmap[hashedData]; !ok {
		network.Kad.hashmap[hashedData] = msg[:]
		reply := []byte("File succesfully stored " + hashedData)
		_, err := resp.servr.WriteToUDP(reply, resp.resp)
		ErrorHandler(err)
	} else {
		reply := []byte(string(msg[:]) + " " + ("File already stored"))
		_, err := resp.servr.WriteToUDP(reply, resp.resp)
		ErrorHandler(err)

	}
	contact := NewContact(network.Contact.ID, network.Contact.Address)

	reply := []byte("Store answer " + network.Contact.ID.String() + " " + network.Contact.Address)
	_, err := resp.servr.WriteToUDP(reply, resp.resp)
	ErrorHandler(err)
	contactArr := make([]Contact, 1)
	contactArr[0] = contact
	return contactArr
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

	_, err := resp.servr.WriteToUDP(ContactToByte(closeContactsArr), resp.resp)
	ErrorHandler(err)
	//closeContactsArr = make([]Contact, 0)
	//closeContactsArr = append(closeContactsArr,))
	returnmsg := make([]Contact, 1)
	returnmsg[0] = NewContact(nil, strconv.Itoa(len(closeContactsArr)))
	return returnmsg
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
	case string(msg[:5]) == "store":
		returnContact = network.HandleStoreMsg(msg[6:], resp)
	case string(msg[:10]) == "FIND_VALUE":
		returnContact = network.HandleFindDataMsg(msg[11:], resp)

	default:
		returnContact = append(returnContact, NewContact(nil, ""))
	}
	return returnContact
}

func (network *Network) rpcHandle(msg data, resp response) []Contact {
	var returnContact []Contact
	switch {
	case msg.rpc == "ping":
		returnContact = network.HandlePingMsg([]byte(msg.data), resp)
		network.Kad.Rtable.AddContact(returnContact[0])
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
	msg := data{}
	for {
		msgbuf := make([]byte, 2048)
		n, resp, err := servr.ReadFromUDP(msgbuf)
		err = json.Unmarshal(msgbuf[:n], &msg)
		//ErrorHandler(err)
		fmt.Println(err)
		Response := &response{
			servr: servr,
			resp:  resp,
		}

		//handledContact := network.msgHandle(msgbuf[:n], *Response)
		fmt.Println("Msg from a friend: ", msg.rpc+" "+msg.data)
		handledContact := network.rpcHandle(msg, *Response)
		fmt.Println("\nResponded with:  ", handledContact)

	}
}
func (network *Network) SendPingMessage(contact *Contact) {
	me := network.Contact
	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	ErrorHandler(err)
	defer connection.Close()
	msg, err := json.Marshal(PingMsg(me))
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

func (network *Network) SendFindDataMessage(hash string, contact *Contact, found chan []Contact, value chan string) {
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
	if string(respmsg[:2]) == "OK" {
		found <- c
		value <- string(respmsg[:n])
	} else {
		c = ByteToContact(respmsg[:n])
		found <- c
		value <- ""
	}
}

func FindDataMsg(hash string) []byte {
	msg := []byte("FIND_VALUE " + hash)
	return msg
}

func (network *Network) HandleFindDataMsg(msg []byte, resp response) []Contact {
	if value, ok := network.Kad.hashmap[string(msg)]; !ok {
		closeContactsArr := network.Kad.Rtable.FindClosestContacts(NewKademliaID(string(msg[:40])), 20)
		_, err := resp.servr.WriteToUDP(ContactToByte(closeContactsArr), resp.resp)
		ErrorHandler(err)
		return closeContactsArr
	} else {
		reply := []byte("OK: " + string(value))
		_, err := resp.servr.WriteToUDP(reply, resp.resp)
		ErrorHandler(err)

		me := make([]Contact, 0)
		me = append(me, *network.Contact)
		return me
	}

}

func (network *Network) IterativeFindNode() []Contact {
	result := make(chan []Contact)
	go network.Kad.LookupContact(*network, result, *network.Contact)
	done := <-result
	fmt.Printf("\nIterativeFindNode done, found %d contacts\n", len(done))
	return done
}

func (network *Network) IterativeStore(data []byte) {
	network.Kad.Store(data, network)
}

func (network *Network) IterativeFindData(hash string) {

	result := network.Kad.LookupData(*network, NewContact(NewKademliaID(hash), ""), hash)
	if result[:2] == "OK" {
		fmt.Println("Value found: " + result[4:])
	} else {
		fmt.Println("Value not found. \nK closest contacts: " + result)
	}

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

func ContactToByte(contactArr []Contact) []byte {

	closeCToByte := make([]byte, 0)
	closeContactsByte := make([]byte, 0)

	for i := 0; i < len(contactArr); i++ {
		closeCToByte = []byte(contactArr[i].ID.String() + " " + contactArr[i].Address + " " + contactArr[i].distance.String() + "\n")
		closeContactsByte = append(closeContactsByte, closeCToByte[:]...)
	}

	return closeContactsByte

}

func (network *Network) SendStoreMessage(contact *Contact, data []byte) {
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
