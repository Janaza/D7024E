package D7024E

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
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
	Rpc      string    `json:"rpc,omitempty"`
	Id       string    `json:"id,omitempty"`
	Ip       string    `json:"ip,omitempty"`
	Contacts []Contact `json:"data,omitempty"`
}

const ID_INDEX = 40

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//creates the content of ping
func PingMsg(contact *Contact) *data {
	msg := &data{
		Rpc: "ping",
		Id:  contact.ID.String(),
		Ip:  contact.Address,
	}
	return msg
}

//creates the content of any message
func createMsg(rpc string, contact *Contact, c []Contact) *data {
	msg := &data{
		Rpc:      rpc,
		Id:       contact.ID.String(),
		Ip:       contact.Address,
		Contacts: c,
	}
	return msg
}

func (network *Network) HandleStoreMsg(msg string, resp response) {
	hashedData := HashData([]byte(msg))
	fmt.Println(hashedData)

	if _, ok := network.Kad.hashmap[hashedData]; !ok {
		network.Kad.hashmap[hashedData] = []byte(msg)
		reply := []byte("File succesfully stored " + hashedData)
		_, err := resp.servr.WriteToUDP(reply, resp.resp)
		ErrorHandler(err)
	} else {
		reply := []byte(msg + " " + ("File already stored"))
		_, err := resp.servr.WriteToUDP(reply, resp.resp)
		ErrorHandler(err)

	}
}

//handles incoming ping msgs
func (network *Network) HandlePingMsg(msg data, r response) {
	contact := NewContact(NewKademliaID(msg.Id), msg.Ip)
	network.Kad.Rtable.AddContact(contact)
	network.SendPongMessage(r)
}

func (network *Network) HandlePongMsg(msg data) {
	contact := NewContact(NewKademliaID(msg.Id), msg.Ip)
	network.Kad.Rtable.AddContact(contact)
	fmt.Println("Added ponger: " + contact.Address)
}

func (network *Network) HandleFindNodeMsg(msg data, r response) {
	Response := &data{
		Contacts: network.Kad.Rtable.FindClosestContacts(NewKademliaID(msg.Id), 20),
	}
	m, err := json.Marshal(Response)
	_, err = r.servr.WriteToUDP(m, r.resp)
	ErrorHandler(err)
	fmt.Println("SENT: my contacts")

}

func (network *Network) HandleFindDataMsg(msg string, r response) {
	if value, ok := network.Kad.hashmap[msg]; !ok {
		closeContactsArr := network.Kad.Rtable.FindClosestContacts(NewKademliaID(msg), 20)
		_, err := r.servr.WriteToUDP(ContactToByte(closeContactsArr), r.resp)
		ErrorHandler(err)
	} else {
		reply := []byte("OK: " + string(value))
		_, err := r.servr.WriteToUDP(reply, r.resp)
		ErrorHandler(err)
	}
}

func (network *Network) rpcHandle(msg data, r response) {
	switch {
	case strings.ToLower(msg.Rpc) == "ping":
		network.HandlePingMsg(msg, r)
	case strings.ToLower(msg.Rpc) == "pong":
		network.HandlePongMsg(msg)
	case strings.ToLower(msg.Rpc) == "find_node":
		network.HandleFindNodeMsg(msg, r)
	case strings.ToLower(msg.Rpc) == "store":
		network.HandleStoreMsg(msg.Id, r)
	case strings.ToLower(msg.Rpc) == "find_value":
		network.HandleFindDataMsg(msg.Id, r)
	default:
		fmt.Println("Unknown RPC: " + msg.Rpc)
	}
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
		msgbuf := make([]byte, 65536)
		n, resp, err := servr.ReadFromUDP(msgbuf)
		ErrorHandler(err)
		json.Unmarshal(msgbuf[:n], &msg)
		Response := &response{
			servr: servr,
			resp:  resp,
		}
		fmt.Println("GOT: ", msg)
		go network.rpcHandle(msg, *Response)

	}
}
func (network *Network) SendPingMessage(contact *Contact) {
	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	ErrorHandler(err)
	defer connection.Close()
	msg, err := json.Marshal(createMsg("ping", network.Contact, nil))
	_, err = connection.Write(msg)
	ErrorHandler(err)
	fmt.Println("SENT: " + string(msg))
	respmsg := make([]byte, 65536)
	n, err := connection.Read(respmsg)
	data := data{}
	err = json.Unmarshal(respmsg[:n], &data)
	ErrorHandler(err)
	network.rpcHandle(data, response{})
}

func (network *Network) SendPongMessage(r response) {
	msg, err := json.Marshal(createMsg("pong", network.Contact, nil))
	ErrorHandler(err)
	_, err = r.servr.WriteToUDP(msg, r.resp)
	ErrorHandler(err)
	fmt.Println("SENT: ", "pong ", network.Contact)
}

func (network *Network) SendFindContactMessage(contact *Contact, found chan []Contact, sl *Shortlist) {
	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	connection.SetDeadline(time.Now().Add(50 * time.Millisecond))
	ErrorHandler(err)
	defer connection.Close()
	msg, err := json.Marshal(createMsg("find_node", contact, nil))
	ErrorHandler(err)
	_, err = connection.Write(msg)
	fmt.Println("SENT: " + string(msg) + " to: " + contact.Address)
	respmsg := make([]byte, 65536)
	data := data{}
	connection.SetReadDeadline(time.Now().Add(50 * time.Millisecond))
	var c []Contact
	for {
		n, err := connection.Read(respmsg)
		if err != nil {
			if e, ok := err.(net.Error); !ok && !e.Timeout() {
				ErrorHandler(e)
				c = make([]Contact, 0)
				break
			}
			fmt.Println("Node offline!")
			sl.removeContact(*contact)
			c = make([]Contact, 0)
			break
		}
		err = json.Unmarshal(respmsg[:n], &data)
		ErrorHandler(err)
		c = data.Contacts
		break
	}
	found <- c
}

func (network *Network) SendFindDataMessage(hash string, contact *Contact, found chan []Contact, value chan string) {
	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	ErrorHandler(err)
	defer connection.Close()
	d := &data{Id: hash, Rpc: "find_value"}
	msg, err := json.Marshal(d)
	_, err = connection.Write(msg)
	ErrorHandler(err)

	respmsg := make([]byte, 65536)
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

func (network *Network) SendStoreMessage(contact *Contact, b []byte) {
	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	ErrorHandler(err)
	defer connection.Close()
	Response := &data{Id: string(b), Rpc: "store"}
	m, err := json.Marshal(Response)
	_, err = connection.Write(m)
	ErrorHandler(err)
	respmsg := make([]byte, 65536)
	n, err := connection.Read(respmsg)
	ErrorHandler(err)
	fmt.Println(string(respmsg[:n]))
}

func (network *Network) IterativeFindNode() []Contact {
	result := make(chan []Contact)
	go network.Kad.LookupContact(*network, result, *network.Contact)
	done := <-result
	fmt.Printf("\nIterativeFindNode done, found %d contacts \n", len(done))
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
