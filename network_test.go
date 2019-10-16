package D7024E

import (
	"reflect"
	"testing"
)

/*func TestInitNode(t *testing.T) {
	contact := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"), "localhost:8000")

	testNetwork := InitNode("localhost",8000, &contact)
	fmt.Println(testNetwork.Contact.Address)
	testNetwork.Listen(contact, 8000)
}*/

func TestInitNode(t *testing.T) {
	newCont := Contact{
		ID:       NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
		Address:  "1",
		distance: NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
	}
	newKad := Kademlia{
		Rtable:  NewRoutingTable(newCont),
		hashmap: make(map[string][]byte),
	}
	newNet := Network{
		Contact: &newCont,
		Kad:     &newKad,
	}

	tests := []struct {
		args *Contact
		want *Network
	}{
		{&newCont, &newNet},
	}
	for _, tt := range tests {
		t.Run(tt.args.Address, func(t *testing.T) {
			if got := InitNode(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
//initbootstrap port 1, och myport 2
func TestInitJoin(t *testing.T) {
	type args struct {
		myport string
		bIP    string
	}
	tests := []struct {
		name string
		args args
		want *Network
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitJoin(tt.args.myport, tt.args.bIP); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitBootstrap(t *testing.T) {

	newCont := Contact{
		ID:       NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
		Address:  "localhost:8000",
		distance: NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
	}
	newKad := Kademlia{
		Rtable:  NewRoutingTable(newCont),
		hashmap: make(map[string][]byte),
	}
	newNet := Network{
		Contact: &newCont,
		Kad:     &newKad,
	}
	tests := []struct {
		args string
		want *Network
	}{
		{"8000",&newNet},
	}
	for _, tt := range tests {
		t.Run("Init bootstrap", func(t *testing.T) {
			if got := InitBootstrap(tt.args); !reflect.DeepEqual(got.Contact.Address, tt.want.Contact.Address) {
				t.Errorf("InitBootstrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorHandler(t *testing.T) {
	tests := []struct {
		args error
	}{
		{nil},
	}
	for _, tt := range tests {
		t.Run("No errors found", func(t *testing.T) {
			ErrorHandler(tt.args)
		})
	}
}

func Test_createMsg(t *testing.T) {
	type args struct {
		rpc     string
		contact *Contact
		c       []Contact
	}
	tests := []struct {
		name string
		args args
		want *data
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createMsg(tt.args.rpc, tt.args.contact, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetwork_updateBuckets(t *testing.T) {
	type fields struct {
		Contact *Contact
		Kad     *Kademlia
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := &Network{
				Contact: tt.fields.Contact,
				Kad:     tt.fields.Kad,
			}
			network.updateBuckets()
		})
	}
}
