package d7024e

type Network struct {
}

func Listen(ip string, port int) {
	// TODO
}

func (network *Network) SendPingMessage(contact *Contact) {
	// TODO PING
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO FIND_NODE
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO FIND_VALUE
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO STORE
}
