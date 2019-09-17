package D7024E

type Kademlia struct {
	Rtable *RoutingTable
}

func InitKad(me Contact) *Kademlia {
	node := &Kademlia{
		Rtable: NewRoutingTable(me),
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
