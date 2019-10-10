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

/*func TestHandleStore(t *testing.T) {
	me := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"), "localhost:8000")
	network := &Network{
		IP:      "",
		Port:    0,
		Contact: &me,
		Kad:     nil,
	}
	resp := response{
		servr: nil,
		resp:  nil,
	}
	msg := []byte("hej")
	returnCon := network.HandleStoreMsg(msg, resp)
	fmt.Println(returnCon[0])

}
*/
/*func HandlePongMsg(msg []byte) Contact {
	contactID := NewKademliaID(string(msg[:ID_INDEX]))
	ipAndPort := ipToString(msg)
	ipAndPortstring := strings.Split(ipAndPort, ":")
	ip, port := ipAndPortstring[0], ipAndPortstring[1]
	contactAdrs := ip
	contactPort := port
	contact := NewContact(contactID, contactAdrs+":"+contactPort)
	return contact
}*/
func TestHandlePongMsg(t *testing.T) {
	newKadID := NewRandomKademliaID()
	const testAdr = "localhost:8000"
	handeledPongMsg := HandlePongMsg([]byte(newKadID.String() +" localhost:8000"))
	fmt.Print("The actual contact: ")
	fmt.Print(handeledPongMsg)
	if handeledPongMsg.ID.String() == newKadID.String() && handeledPongMsg.Address== testAdr {
		fmt.Println("Pong Contact is handled correctly")

	}else{
		fmt.Println(handeledPongMsg.ID.String())
		fmt.Println("Pong Contact is incorrect")
	}
	fmt.Println("Expected " + newKadID.String() +" " + testAdr)
	fmt.Println("Acctual result: " + handeledPongMsg.ID.String() +" "+testAdr )


}
