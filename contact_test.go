package D7024E

import (
	"reflect"
	"testing"
)

func TestNewContact(t *testing.T) {
	type args struct {
		id      *KademliaID
		address string
	}
	tests := []struct {
		name string
		args args
		want Contact
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContact(tt.args.id, tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContact_CalcDistance(t *testing.T) {
	type fields struct {
		ID       *KademliaID
		Address  string
		distance *KademliaID
	}
	type args struct {
		target *KademliaID
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
			contact := &Contact{
				ID:       tt.fields.ID,
				Address:  tt.fields.Address,
				distance: tt.fields.distance,
			}
			contact.CalcDistance(tt.args.target)
		})
	}
}
