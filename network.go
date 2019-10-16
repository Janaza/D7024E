package D7024E

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
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
//1. Grabs ip from eth0 and generates a new nodeID
//2. Inserts some known node by pinging a given bootstrap ip
//3. Finds new nodes with IterativeFindNode on nodeID (me)
//4. Refreshes all buckets further away than its closest neighbor.
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
	n.updateBuckets()
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

//Holds a UDP connection
type response struct {
	servr *net.UDPConn
	resp  *net.UDPAddr
}

//Holds packet information as JSON
type data struct {
	Rpc      string    `json:"rpc,omitempty"`
	Id       string    `json:"id,omitempty"`
	Ip       string    `json:"ip,omitempty"`
	Contacts []Contact `json:"data,omitempty"`
}

const ID_INDEX = 40

//log any errors
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

//Calls functions by RPC value
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

//Starts UDP listener
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

//Helper function to call LookupContact()
func (network *Network) IterativeFindNode(target *Contact) []Contact {
	result := make(chan []Contact)
	go network.Kad.LookupContact(network.Contact, result, *target)
	done := <-result
	fmt.Printf("\nIterativeFindNode done, found %d contacts \n", len(done))
	return done
}

//Helper function to call Store()
func (network *Network) IterativeStore(data []byte) {
	network.Kad.Store(data, network.Contact)
}

//Helper function to call LookupData()
func (network *Network) IterativeFindData(hash string) {
	result := network.Kad.LookupData(network.Contact, NewContact(NewKademliaID(hash), ""), hash)
	if result[:2] == "OK" {
		fmt.Println("Value found: " + result[4:])
	} else {
		fmt.Println("Value not found. \nK closest contacts: " + result)
	}

}

//refreshes all buckets further away than closest neighbor
func (network *Network) updateBuckets() {
	//Loop over any populated bucket
	bucketPopulated := false
	//String reprensatation of bits in kadid
	sArr := make([]string, IDLength)
	//Loop over each byte
	for i := range [IDLength]int{} {
		//byte to bits
		bits := strconv.FormatInt(int64(network.Contact.ID[i]), 2)
		for j := len(bits); j < 8; j++ {
			bits = "0" + bits
		}
		sArr[i] = bits
	}
	//Start at LSByte
	for i := IDLength - 1; i >= 0; i-- {
		//Flip each bit at current byte
		for j := 7; j >= 0; j-- {
			fliped := sArr
			if string(fliped[i][j]) == "1" {
				fliped[i] = fliped[i][:j] + string("0") + fliped[i][j+1:]
			} else {
				fliped[i] = fliped[i][:j] + string("1") + fliped[i][j+1:]
			}

			toBytearr := BitsToKademliaID(fliped)
			//Update bucket at newID
			newID := NewKademliaID(hex.EncodeToString(toBytearr[:IDLength]))
			c := NewContact(newID, "")
			if bucketPopulated == true || network.Kad.Rtable.buckets[network.Kad.Rtable.GetBucketIndex(newID)].Len() != 0 {
				arr := network.IterativeFindNode(&c)
				for _, c := range arr {
					network.Kad.Rtable.AddContact(c)
				}
				bucketPopulated = true
			}
		}
	}
}
