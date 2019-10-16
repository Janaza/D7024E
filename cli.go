package D7024E

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func (network *Network) CliHelper(input io.Reader) {
	for {
		network.Cli(input)
	}
}
func (network *Network) Cli(input io.Reader) {

	cli := bufio.NewScanner(input)
	fmt.Printf("Command: \n")
	cli.Scan()
	text := cli.Text()
	fmt.Println(text)

	switch {
	case strings.Contains(text, "PUT "):
		storeData := []byte(text[4:])
		fmt.Println("Storing data on other nodes")
		network.IterativeStore(storeData)

	case strings.Contains(text, "GET "):
		hashData := text[4:]
		fmt.Println("Fetching data...")
		if len(hashData) == 40 {
			network.IterativeFindData(hashData)
		} else {
			fmt.Println("The length of hash is wrong.")
		}

	case text == "EXIT":
		fmt.Println("Node is shutting down in 3 seconds...")
		time.Sleep(3 * time.Second)
		os.Exit(0)

	//Must ping an address
	case strings.Contains(text, "PING "):
		node := NewContact(nil, text[5:])
		msg, err := SendPingMessage(&node, network.Contact)
		if err == nil {
			network.rpcHandle(msg, response{})
		}

	//Must store a 3 characters to a given IP
	case strings.Contains(text, "STORE "):
		storeData := []byte(text[6:9])
		node := NewContact(nil, text[10:])
		SendStoreMessage(&node, storeData)

	case text == "CONTACTS":
		for _, i := range network.Kad.Rtable.FindClosestContacts(NewKademliaID("0000000000000000000000000000000000000000"), 160) {
			fmt.Println(i.Address)
		}

	default:
		fmt.Sprintln("CLI not recognized")
	}

}
