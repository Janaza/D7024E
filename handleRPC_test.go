package D7024E

import "testing"

var newCont = Contact{
ID:       NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
Address:  "localhost:8000",
distance: NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
}

var newKad = Kademlia{
Rtable:  NewRoutingTable(newCont),
hashmap: make(map[string][]byte),
}
var newNet = Network{
Contact: &newCont,
Kad:     &newKad,
}

func TestNetwork_HandleStoreMsg(t *testing.T) {
	type fields struct {
		Contact *Contact
		Kad     *Kademlia
	}
	type args struct {
		msg  string
		resp response
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := &Network{
				Contact: tt.fields.Contact,
				Kad:     tt.fields.Kad,
			}
			network.HandleStoreMsg(tt.args.msg, tt.args.resp)
		})
	}
}
