package D7024E

import (
	"reflect"
	"testing"
)

func TestInitKad(t *testing.T) {
	type args struct {
		me Contact
	}
	tests := []struct {
		name string
		args args
		want *Kademlia
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitKad(tt.args.me); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitKad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKademlia_LookupContact(t *testing.T) {
	type fields struct {
		Rtable  *RoutingTable
		hashmap map[string][]byte
	}
	type args struct {
		me     *Contact
		result chan []Contact
		target Contact
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
			kademlia := &Kademlia{
				Rtable:  tt.fields.Rtable,
				hashmap: tt.fields.hashmap,
			}
			kademlia.LookupContact(tt.args.me, tt.args.result, tt.args.target)
		})
	}
}

func TestKademlia_LookupData(t *testing.T) {
	type fields struct {
		Rtable  *RoutingTable
		hashmap map[string][]byte
	}
	type args struct {
		me     *Contact
		target Contact
		hash   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kademlia := &Kademlia{
				Rtable:  tt.fields.Rtable,
				hashmap: tt.fields.hashmap,
			}
			if got := kademlia.LookupData(tt.args.me, tt.args.target, tt.args.hash); got != tt.want {
				t.Errorf("Kademlia.LookupData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKademlia_Store(t *testing.T) {
	type fields struct {
		Rtable  *RoutingTable
		hashmap map[string][]byte
	}
	type args struct {
		data []byte
		me   *Contact
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
			kademlia := &Kademlia{
				Rtable:  tt.fields.Rtable,
				hashmap: tt.fields.hashmap,
			}
			kademlia.Store(tt.args.data, tt.args.me)
		})
	}
}

func TestShortlist_insert(t *testing.T) {
	type fields struct {
		ls []Contact
		v  map[string]bool
	}
	type args struct {
		v bool
		c Contact
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Contact
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := &Shortlist{
				ls: tt.fields.ls,
				v:  tt.fields.v,
			}
			if got := sl.insert(tt.args.v, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shortlist.insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortlist_removeContact(t *testing.T) {
	type fields struct {
		ls []Contact
		v  map[string]bool
	}
	type args struct {
		c Contact
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
			sl := &Shortlist{
				ls: tt.fields.ls,
				v:  tt.fields.v,
			}
			sl.removeContact(tt.args.c)
		})
	}
}
