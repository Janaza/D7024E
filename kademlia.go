package D7024E

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type Kademlia struct {
	Rtable *RoutingTable
	data []dataStruct
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
	hashedData := HashData(data)
	for _, dataFile := range kademlia.data{
		if dataFile.Hash == hashedData{
			fmt.Println("File already stored")
			return
		} else{
			fmt.Println("Storing " + string(data) + " with hash "+ hashedData)
			dataFile := dataStruct{hashedData, data, }
			kademlia.data = append(kademlia.data, dataFile)
		}
	}
}

func HashData(data []byte) string{
	hashedData := sha1.Sum(data)
	hashedStringdata := hex.EncodeToString(hashedData[0:20])
	return hashedStringdata
}