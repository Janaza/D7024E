package D7024E

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func SendPingMessage(contact *Contact, me *Contact) (data, error) {
	RemoteAddress, err := net.ResolveUDPAddr("udp", contact.Address)
	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	ErrorHandler(err)
	defer connection.Close()
	msg, err := json.Marshal(createMsg("ping", me, nil))
	_, err = connection.Write(msg)
	ErrorHandler(err)
	fmt.Println("SENT: " + string(msg))
	data := data{}
	respmsg := make([]byte, 65536)
	connection.SetReadDeadline(time.Now().Add(1 * time.Second))
	for {
		n, err := connection.Read(respmsg)
		if err != nil {
			if e, ok := err.(net.Error); !ok && !e.Timeout() {
				ErrorHandler(e)
				break
			}
			fmt.Println("Node offline!")
			break
		}
		err = json.Unmarshal(respmsg[:n], &data)
		ErrorHandler(err)
		return data, nil
		break
	}
	return data, err
}

func SendPongMessage(r response, me *Contact) {
	msg, err := json.Marshal(createMsg("pong", me, nil))
	ErrorHandler(err)
	_, err = r.servr.WriteToUDP(msg, r.resp)
	ErrorHandler(err)
	fmt.Println("SENT: ", "pong ", me)
}

func SendFindContactMessage(contact *Contact, found chan []Contact, sl *Shortlist) {
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

func SendFindDataMessage(hash string, contact *Contact, found chan []Contact, value chan string) {
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

func SendStoreMessage(contact *Contact, b []byte) {
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
