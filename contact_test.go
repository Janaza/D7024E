package D7024E

import (
	"testing"
)

func Test_NewContact(t *testing.T){

	got := NewContact(NewKademliaID("ffffffff00000000000000000000000000000001"),"localhost:1")
	if got.String() != "contact(\"ffffffff00000000000000000000000000000001\", \"localhost:1\")"{
		t.Errorf("got.String() = %v; want contact(\"ffffffff00000000000000000000000000000001\", \"localhost:1\")", got.String())
	}

}

func TestContact_CalcDistance(t *testing.T) {
	got := NewContact(NewKademliaID("ffffffff00000000000000000000000000000001"), "localhost:1")
	test1 := NewKademliaID("ffffffff00000000000000000000000000000001")
	got.CalcDistance(test1)
	if got.distance.String() != "0000000000000000000000000000000000000000" {
		t.Errorf("got.distance = %v; want 0000000000000000000000000000000000000000", got)
	}
}

func TestContact_Less(t *testing.T) {
	contact1 := NewContact(NewKademliaID("ffffffff00000000000000000000000000000002"), "localhost:1")
	contact2 := NewContact(NewKademliaID("ffffffff00000000000000000000000000000001"), "localhost:2")
	testID := NewKademliaID("ffffffff00000000000000000000000000000003")
	contact1.CalcDistance(testID)
	contact2.CalcDistance(testID)
	got := contact1.Less(&contact2)
	if got == false{
		t.Errorf("got.Less(&contact2) = %v; want true", got)
	}
}

func TestContact_String(t *testing.T) {
	got := NewContact(NewKademliaID("ffffffff00000000000000000000000000000001"), "localhost:2")
	test := "contact(\"ffffffff00000000000000000000000000000001\", \"localhost:2\")"
	if got.String() != test {
		t.Errorf("got.String() = %v; want contact(\"ffffffff00000000000000000000000000000001\", \"localhost:2\")", got)
	}
}

func TestContactCandidates_Append(t *testing.T) {
	var got ContactCandidates
	contactstest := make([]Contact, 0)
	contactstest = append(contactstest, NewContact(NewKademliaID("ffffffff00000000000000000000000000000001"), "localhost:2"))
	got.Append(contactstest)
	if got.GetContacts(1)[0] != contactstest[0]{
		t.Errorf("got.GetContacts(1)[0] = %v; want {ffffffff00000000000000000000000000000001 localhost:2 <nil>}",  got)
	}
}

func TestContactCandidates_GetContacts(t *testing.T) {
	var candidates ContactCandidates
	contactstest := make([]Contact, 0)
	contactstest = append(contactstest, NewContact(NewKademliaID("ffffffff00000000000000000000000000000001"), "localhost:2"))
	contactstest = append(contactstest, NewContact(NewKademliaID("ffffffff00000000000000000000000000000002"), "localhost:3"))
	contactstest = append(contactstest, NewContact(NewKademliaID("ffffffff00000000000000000000000000000003"), "localhost:4"))
	contactstest = append(contactstest, NewContact(NewKademliaID("ffffffff00000000000000000000000000000004"), "localhost:5"))
	candidates.Append(contactstest)

	got := candidates.GetContacts(4)
	for i := range got{
		if got[i] != contactstest[i]{
			t.Errorf("got[%v] = %v; want %v", i,  got[i], contactstest[i])
		}
	}
}

/*
// Sort the Contacts in ContactCandidates
func (candidates *ContactCandidates) Sort() {
	sort.Sort(candidates)
}


func TestContactCandidates_Sort(t *testing.T) {
	var candidates1 ContactCandidates
	var candidates2 ContactCandidates

	contactstest1 := make([]Contact, 0)
	contactstest1 = append(contactstest1, NewContact(NewKademliaID("ffffffff00000000000000000000000000000002"), "localhost:2"))
	contactstest1 = append(contactstest1, NewContact(NewKademliaID("ffffffff00000000000000000000000000000001"), "localhost:1"))
	candidates1.Append(contactstest1)
	candidates1.Sort()
	contactstest2 := make([]Contact, 0)
	contactstest2 = append(contactstest2, NewContact(NewKademliaID("ffffffff00000000000000000000000000000001"), "localhost:1"))
	contactstest2 = append(contactstest2, NewContact(NewKademliaID("ffffffff00000000000000000000000000000002"), "localhost:2"))
	candidates2.Append(contactstest2)
	got := candidates1.GetContacts(2)
	want := candidates2.GetContacts(2)
	for i := range got{
		if got[i].String() != want[i].String(){
			t.Errorf("got[%v] = %v; want %v", i,  got[i], want[i])
		}
	}
}
*/


func TestContactCandidates_Len(t *testing.T) {
	var got ContactCandidates
	contactstest := make([]Contact, 0)
	contactstest = append(contactstest, NewContact(NewKademliaID("ffffffff00000000000000000000000000000001"), "localhost:2"))
	contactstest = append(contactstest, NewContact(NewKademliaID("ffffffff00000000000000000000000000000002"), "localhost:3"))
	got.Append(contactstest)
	if got.Len() != 2 {
		t.Errorf("got.Len() = %v; want 2", got.Len())
	}
}

func TestContactCandidates_Swap(t *testing.T) {

	var candidates1 ContactCandidates
	var candidates2 ContactCandidates

	contactstest1 := make([]Contact, 0)
	contactstest1 = append(contactstest1, NewContact(NewKademliaID("ffffffff00000000000000000000000000000001"), "localhost:2"))
	contactstest1 = append(contactstest1, NewContact(NewKademliaID("ffffffff00000000000000000000000000000002"), "localhost:3"))
	candidates1.Append(contactstest1)
	candidates1.Swap(0, 1)
	contactstest2 := make([]Contact, 0)
	contactstest2 = append(contactstest2, NewContact(NewKademliaID("ffffffff00000000000000000000000000000002"), "localhost:3"))
	contactstest2 = append(contactstest2, NewContact(NewKademliaID("ffffffff00000000000000000000000000000001"), "localhost:2"))
	candidates2.Append(contactstest2)
	got := candidates1.GetContacts(2)
	want := candidates2.GetContacts(2)

	for i := range got{
		if got[i].String() != want[i].String(){
			t.Errorf("got[%v] = %v; want %v", i,  got[i], want[i])
		}
	}
}

/*
// Less returns true if the Contact at index i is smaller than
// the Contact at index j
func (candidates *ContactCandidates) Less(i, j int) bool {
	return candidates.contacts[i].Less(&candidates.contacts[j])
}


func TestContactCandidates_Less(t *testing.T) {

	var candidates ContactCandidates

	contactstest := make([]Contact, 0)
	contactstest = append(contactstest, NewContact(NewKademliaID("ffffffff00000000000000000000000000000001"), "localhost:2"))
	contactstest = append(contactstest, NewContact(NewKademliaID("ffffffff00000000000000000000000000000009"), "localhost:3"))
	candidates.Append(contactstest)
	if candidates.Less(0,1) == false{
		t.Errorf("got = %v; want true",  candidates.Less(0,1))
	}
}

*/