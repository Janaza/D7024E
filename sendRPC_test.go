package D7024E

import (
	"reflect"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func Test_SendRPC(t *testing.T) {
	me := NewContact(NewRandomKademliaID(), "localhost:2")
	bootstrap := InitBootstrap("11")
	fails, err := SendPingMessage(bootstrap.Contact, &me)
	fwant := data{}
	t.Run("SendPingMessage test", func(t *testing.T) {
		if !reflect.DeepEqual(fails, fwant) || err != nil {
			t.Errorf("got = %v, want = %v", fails, fwant)
		}

	})
	final := make(chan []Contact)
	sl := &Shortlist{
		ls: make([]Contact, 0),
		v:  make(map[string]bool),
	}
	go SendFindContactMessage(bootstrap.Contact, final, sl, &me)
	g := <-final
	f := make([]Contact, 0)
	//f = append(f, me)

	t.Run("SendFindContactMessage test", func(t *testing.T) {
		if !reflect.DeepEqual(g, f) {
			t.Errorf("Should fail! got = %v, want = %v", g, f)
		}

	})
	go bootstrap.Listen(*bootstrap.Contact)
	SendPingMessage(bootstrap.Contact, &me)
	r := make(chan []Contact)
	go SendFindContactMessage(bootstrap.Contact, r, sl, &me)
	got := <-r
	want := []Contact{}
	want = append(want, me)

	t.Run("SendFindContactMessage test", func(t *testing.T) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got = %v, want = %v", got, want)
		}

	})
	SendStoreMessage(bootstrap.Contact, []byte("123"))
}
