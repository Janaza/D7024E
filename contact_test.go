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

func TestContact_Less(t *testing.T) {
	type fields struct {
		ID       *KademliaID
		Address  string
		distance *KademliaID
	}
	type args struct {
		otherContact *Contact
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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
			if got := contact.Less(tt.args.otherContact); got != tt.want {
				t.Errorf("Contact.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContact_String(t *testing.T) {
	type fields struct {
		ID       *KademliaID
		Address  string
		distance *KademliaID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
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
			if got := contact.String(); got != tt.want {
				t.Errorf("Contact.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContactCandidates_Append(t *testing.T) {
	type fields struct {
		contacts []Contact
	}
	type args struct {
		contacts []Contact
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
			candidates := &ContactCandidates{
				contacts: tt.fields.contacts,
			}
			candidates.Append(tt.args.contacts)
		})
	}
}

func TestContactCandidates_GetContacts(t *testing.T) {
	type fields struct {
		contacts []Contact
	}
	type args struct {
		count int
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
			candidates := &ContactCandidates{
				contacts: tt.fields.contacts,
			}
			if got := candidates.GetContacts(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContactCandidates.GetContacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContactCandidates_Sort(t *testing.T) {
	type fields struct {
		contacts []Contact
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			candidates := &ContactCandidates{
				contacts: tt.fields.contacts,
			}
			candidates.Sort()
		})
	}
}

func TestContactCandidates_Len(t *testing.T) {
	type fields struct {
		contacts []Contact
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			candidates := &ContactCandidates{
				contacts: tt.fields.contacts,
			}
			if got := candidates.Len(); got != tt.want {
				t.Errorf("ContactCandidates.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContactCandidates_Swap(t *testing.T) {
	type fields struct {
		contacts []Contact
	}
	type args struct {
		i int
		j int
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
			candidates := &ContactCandidates{
				contacts: tt.fields.contacts,
			}
			candidates.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestContactCandidates_Less(t *testing.T) {
	type fields struct {
		contacts []Contact
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			candidates := &ContactCandidates{
				contacts: tt.fields.contacts,
			}
			if got := candidates.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("ContactCandidates.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
