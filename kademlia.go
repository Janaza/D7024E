package d7024e

type Kademlia struct {
	Rtable  *RoutingTable
	Network *Network
}

func InitNode(me Contact, port int) *Kademlia {
	node := &Kademlia{
		Rtable:  NewRoutingTable(me),
		Network: InitNetwork(me.Address, port, me),
	}
	return node
}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO iterativeFindNode

}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO iterativeFindValue

}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO iterativeStore
}
