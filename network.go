package D7024E

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
)

type Network struct {
	Contact *Contact
	Kad     *Kademlia
}

func InitNode(me *Contact) *Network {
	network := &Network{
		Contact: me,
		Kad:     InitKad(*me),
	}
	fmt.Println("I am: ", me.ID, me.Address+"\n")
	return network
}

//InitJoin joins a network with given IP address
//Grabs ip from eth0
//Pings given bootstrap ip and finds new nodes with IterativeFindNode
func InitJoin(myport string, bIP string) *Network {
	myip, err := Eth0IP()
	ErrorHandler(err)

	bContact := NewContact(nil, bIP)
	me := NewContact(NewRandomKademliaID(), myip+":"+myport)
	n := InitNode(&me)
	msg, err := SendPingMessage(&bContact, &me)
	if err == nil {
		n.rpcHandle(msg, response{})
	}
	val := n.IterativeFindNode(&me)
	for _, c := range val {
		n.Kad.Rtable.AddContact(c)
	}
	//n.Kad.Refresh()
	return n
}

//InitBootstrap inits a new node with no known nodes
func InitBootstrap(myport string) *Network {
	myip, err := Eth0IP()
	ErrorHandler(err)
	me := NewContact(NewRandomKademliaID(), myip+":"+myport)
	n := InitNode(&me)
	return n
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
func (network *Network) Listen(contact Contact) {
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

func (network *Network) IterativeFindNode(target *Contact) []Contact {
	result := make(chan []Contact)
	go network.Kad.LookupContact(network.Contact, result, *target)
	done := <-result
	fmt.Printf("\nIterativeFindNode done, found %d contacts \n", len(done))
	return done
}

func (network *Network) IterativeStore(data []byte) {
	network.Kad.Store(data, network.Contact)
}

func (network *Network) IterativeFindData(hash string) {

	result := network.Kad.LookupData(network.Contact, NewContact(NewKademliaID(hash), ""), hash)
	if result[:2] == "OK" {
		fmt.Println("Value found: " + result[4:])
	} else {
		fmt.Println("Value not found. \nK closest contacts: " + result)
	}

}
