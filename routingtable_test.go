package D7024E

import (
	"reflect"
	"testing"
)

func TestNewRoutingTable(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	g := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	fails := NewRoutingTable(NewContact(NewKademliaID("0000000000000000000000000000000000000000"), "localhost:8000"))
	t.Run("NewRoutingTable test", func(t *testing.T) {
		if !reflect.DeepEqual(g, rt) {
			t.Errorf("Routingtable = %v, want = %v", g, rt)
		}
		if reflect.DeepEqual(fails, rt) {
			t.Errorf("Routingtable = want %v, = %v", g, rt)
		}
	})
}

func TestAddContact(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("0000000000000000000000000000000000000000"), "localhost:8000"))
	c := &Contact{
		ID:       NewKademliaID("0000000000000000000000000000000000000001"),
		Address:  "localhost:8000",
		distance: NewKademliaID("0"),
	}
	rt.AddContact(*c)

	for elt := rt.buckets[159].list.Front(); elt != nil; elt = elt.Next() {
		contact := elt.Value.(Contact)
		t.Run("AddContact test", func(t *testing.T) {
			if got := contact; !reflect.DeepEqual(got.ID, c.ID) {
				t.Errorf("cArr = %v, want = %v", got.ID, c.ID)
			}
		})
	}

}

func TestFindClosestContacts(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	var want []Contact
	want = append(want, NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"),
		NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"),
		NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"), "localhost:8001"))

	cArr := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := 0; i < len(cArr); i++ {
		w := want[i].ID.String()
		t.Run("FindClosestContacts test", func(t *testing.T) {
			if got := cArr[i].ID.String(); got != w {
				t.Errorf("cArr = %v, want = %v", got, w)
			}
		})
	}
}

func TestGetBucketByID(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("0000000000000000000000000000000000000000"), "localhost:8000"))
	c := NewKademliaID("0000000000000000000000000000000000000000")
	want := rt.buckets[159]
	fails := rt.buckets[0]

	t.Run("GetBucketByID test", func(t *testing.T) {
		if got := rt.GetBucketByID(c); !(got == want) {
			t.Errorf("GetBucketByID = %v, want = %v", got, want)
		}
		if got := rt.GetBucketByID(c); got == fails {
			t.Errorf("Should differ! GetBucketByID = %v, fails = %v", got, fails)
		}
	})
}

func TestGetBucketIndex(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("0000000000000000000000000000000000000000"), "localhost:8000"))
	c := NewKademliaID("0000000000000000000000000000000000000001")
	want := 159
	fails := 1

	t.Run("GetBucketIndex test", func(t *testing.T) {
		if got := rt.GetBucketIndex(c); !(got == want) {
			t.Errorf("GetBucketIndex = %v, want = %v", got, want)
		}
		if got := rt.GetBucketIndex(c); got == fails {
			t.Errorf("Should differ! GetBucketIndex = %v, fails = %v", got, fails)
		}
	})
}
