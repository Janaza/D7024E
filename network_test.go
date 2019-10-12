package D7024E

import (
	"fmt"
	"testing"
)

func TestPingMsg(t *testing.T) {
	contact := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"), "localhost:8000")
	testMsg := PingMsg(&contact)
	if string(testMsg) == "ping ffffffffffffffffffffffffffffffffffffffff localhost:8000"{
		fmt.Println("PingMsg is working correctly")

	} else{
		fmt.Println("Error in PingMsg")
	}
}
func TestFindNodeMsg(t *testing.T) {
	contact := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"), "localhost:8000")
	testMsg := FindNodeMsg(&contact)
	if string(testMsg) == "FIND_NODE ffffffffffffffffffffffffffffffffffffffff"{
		fmt.Println("FindNodeMsg is working correctly")
	} else{
		fmt.Println("Error in FindNodeMsg")
	}
}
func TestStoreMsg(t *testing.T) {
	data := []byte("123")
	testMsg := StoreMsg(data)
	if string(testMsg) == "store 123"{
		fmt.Println("StoreMsg is working correctly")
	} else{
		fmt.Println("Error in StoreMsg")
	}
}
func TestHandlePongMsg(t *testing.T) {
	msg := []byte("ffffffffffffffffffffffffffffffffffffaaaa localhost:8000")
	pongContact := HandlePongMsg(msg)
	if pongContact.ID.String() == "ffffffffffffffffffffffffffffffffffffaaaa" && pongContact.Address == "localhost:8000"{
		fmt.Println("PongMsg handled correctly")
	}else{
		fmt.Println("Error in PongMessageHandler")
	}

}

/*func TestInitNode(t *testing.T) {
	contact := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"), "localhost:8000")

	testNetwork := InitNode("localhost",8000, &contact)
	fmt.Println(testNetwork.Contact.Address)
	testNetwork.Listen(contact, 8000)
}*/