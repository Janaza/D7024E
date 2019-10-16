package D7024E

import (
	"bytes"
	"testing"
)

func Test_Cli(t *testing.T) {
	me := NewContact(NewRandomKademliaID(), "localhost:20")
	n := InitNode(&me)
	bootstrap := InitBootstrap("10")
	go bootstrap.Listen(*bootstrap.Contact)
	var stdin bytes.Buffer

	tests := []string{
		"PING localhost:10",
		"PUT 123",
		"GET 40bd001563085fc35165329ea1ff5c5ecbdbbeef",
		"STORE 123 localhost:10",
		"CONTACTS",
	}
	for _, s := range tests {
		stdin.Write([]byte(s))
		n.Cli(&stdin)
	}
}
