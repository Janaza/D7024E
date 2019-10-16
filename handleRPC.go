package D7024E

import (
	"encoding/json"
	"fmt"
)

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
	SendPongMessage(r, network.Contact)
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
