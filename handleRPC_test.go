package D7024E

import "testing"

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
		// TODO: Add test cases.
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
