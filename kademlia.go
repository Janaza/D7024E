package D7024E

import (
	"crypto/sha1"
	"encoding/hex"
)

type Kademlia struct {
	Rtable *RoutingTable
	data []dataStruct
	net	Network
}
type dataStruct struct {
	Hash	string
	Value	[]byte
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
	storeContacts := kademlia.iterativeFindNode()
	for i := 0; i < len(storeContacts) && i < 20; i++ {
		kademlia.net.SendStoreMessage(storeContacts[i].contact, data)
	}
}

func HashData(data []byte) string{
	hashedData := sha1.Sum(data)
	hashedStringdata := hex.EncodeToString(hashedData[0:20])
	return hashedStringdata
}