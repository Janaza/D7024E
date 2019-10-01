package D7024E

import (
	"fmt"
	"testing"
)

func TestRoutingTable(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}
}

func TestNetwork_HandleStoreMsg(t *testing.T) {

}
func TestStoreMsg(t *testing.T) {
	storeVal := []byte("h")
	svar := HashData(storeVal)
	fmt.Println(svar)
	contact := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"), "localhost:8000")
	testMsg := StoreMsg(&contact, storeVal)
	fmt.Println(string(testMsg))

}
func TestPingMsg(t *testing.T) {
	contact := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"), "localhost:8000")
	testMsg := PingMsg(&contact)
	if string(testMsg) == "ping ffffffffffffffffffffffffffffffffffffffff localhost:8000"{
		fmt.Println("PingMsg is working correctly")

	} else{
		fmt.Println("Error in PingMsg")
	}

}


func TestHashFunc(t *testing.T){
	hej := []byte("h")
	svar := HashData(hej)
	fmt.Println(svar)
}
