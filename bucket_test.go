package D7024E

import (
	"container/list"
	"reflect"
	"testing"
)

func Test_bucket_AddContact(t *testing.T) {
	type fields struct {
		list *list.List
	}
	type args struct {
		contact Contact
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
			bucket := &bucket{
				list: tt.fields.list,
			}
			bucket.AddContact(tt.args.contact)
		})
	}
}

func Test_bucket_GetContactAndCalcDistance(t *testing.T) {
	type fields struct {
		list *list.List
	}
	type args struct {
		target *KademliaID
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
			bucket := &bucket{
				list: tt.fields.list,
			}
			if got := bucket.GetContactAndCalcDistance(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bucket.GetContactAndCalcDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitsToKademliaID(t *testing.T) {
	type args struct {
		bArr []string
	}
	tests := []struct {
		name string
		args args
		want [IDLength]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BitsToKademliaID(tt.args.bArr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BitsToKademliaID() = %v, want %v", got, tt.want)
			}
		})
	}
}
