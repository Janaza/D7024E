package D7024E

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
)

type Kademlia struct {
	Rtable *RoutingTable
	net	Network
	data   []dataStruct
}
type dataStruct struct {
	Hash  string
	Value []byte
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
	var shortlist []Contact
	doublet := make(map[string]bool)
	for _, mine := range myClosest {
		shortlist = append(shortlist, mine)
		doublet[mine.ID.String()] = true
	}

	runningRoutines := 0
	for runningRoutines < 3 && len(shortlist) > 1 {
		go network.SendFindContactMessage(&shortlist[runningRoutines], found)
		x = <-found
		runningRoutines++
	}
	if len(shortlist) == 1 {
		runningRoutines++
		go network.SendFindContactMessage(&shortlist[0], found)
		x = <-found

	}

	for runningRoutines > 0 {
		recived := x
		for _, candidate := range recived {
			if !(candidate.Address == network.Contact.Address) && !(candidate.ID == nil) {
				if doublet[candidate.ID.String()] == false {
					doublet[candidate.ID.String()] = true
					candidate.CalcDistance(network.Contact.ID)
					shortlist = append(shortlist, candidate)
				}
			}
		}
		shortlist = qsort(shortlist, target)
		runningRoutines--
		if closestNode.ID.String() != shortlist[0].ID.String() {
			closestNode.ID = shortlist[0].ID
			for i := range shortlist {
				if i >= 3 {
					break
				}
				runningRoutines++
				go network.SendFindContactMessage(&shortlist[i], found)
				x = <-found

			}
		}
	}
	shortlist = qsort(shortlist, target)
	/* //Shows that list is sorted
	for _, c := range shortlist {
		dist := c.ID.CalcDistance(target.ID)
		fmt.Println(dist.String())
	}
	*/
	if len(shortlist) > 20 {
		result <- shortlist[:20]
	} else {
		result <- shortlist
	}
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

func HashData(data []byte) string {
	hashedData := sha1.Sum(data)
	hashedStringdata := hex.EncodeToString(hashedData[0:20])
	return hashedStringdata
}

func qsort(contact []Contact, target Contact) []Contact {
	if len(contact) < 2 {
		return contact
	}

	left, right := 0, len(contact)-1

	pivot := rand.Int() % len(contact)

	contact[pivot], contact[right] = contact[right], contact[pivot]

	for i := range contact {
		dist := contact[i].ID.CalcDistance(target.ID)
		distr := contact[right].ID.CalcDistance(target.ID)
		if dist.Less(distr) {
			contact[left], contact[i] = contact[i], contact[left]
			left++
		}
	}

	contact[left], contact[right] = contact[right], contact[left]

	qsort(contact[:left], target)
	qsort(contact[left+1:], target)

	return contact
}
