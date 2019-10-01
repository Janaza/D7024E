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

func (kademlia *Kademlia) LookupContact(network Network, result chan []Contact, target Contact) {
	alpha := 3
	found := make(chan []Contact)
	var closestNode Contact
	var x []Contact
	myClosest := network.Kad.Rtable.FindClosestContacts(target.ID, alpha)
	closestNode = myClosest[0]
	shortlist := ContactCandidates{}
	doublet := make(map[Contact]bool)
	for _, mine := range myClosest {
		shortlist.contacts = append(shortlist.contacts, mine)
		doublet[mine] = true
	}

	runningRoutines := 0
	for runningRoutines < 3 && len(shortlist.contacts) > 1 {
		go network.SendFindContactMessage(&shortlist.contacts[runningRoutines], found)
		x = <-found
		runningRoutines++
	}
	if len(shortlist.contacts) == 1 {
		runningRoutines++
		go network.SendFindContactMessage(&shortlist.contacts[0], found)
		x = <-found

	}

	for runningRoutines > 0 {
		recived := x
		fmt.Printf("\nrunningRoutines: %d\n", runningRoutines)
		for _, candidate := range recived {
			if doublet[candidate] != true && !(candidate.Address == network.Contact.Address) && !(candidate.ID == nil) {
				doublet[candidate] = true
				shortlist.contacts = append(shortlist.contacts, candidate)
			}
		}

		shortlist.Sort()
		runningRoutines--
		if closestNode != shortlist.contacts[0] {
			closestNode = shortlist.contacts[0]
			for i := range shortlist.contacts {
				if i >= 3 {
					break
				}
				runningRoutines++
				go network.SendFindContactMessage(&shortlist.contacts[i], found)
				x = <-found

			}
		}
	}
	shortlist.Sort()
	if len(shortlist.contacts) > 20 {
		result <- shortlist.contacts[:20]
	} else {
		result <- shortlist.contacts
	}
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO iterativeFindValue

}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO iterativeStore
}

func (kademlia *Kademlia) HashData(data []byte) string{
	hashedData := sha1.Sum(data)
	hashedStringdata := hex.EncodeToString(hashedData[0:20])
	return hashedStringdata
}