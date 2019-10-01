package D7024E

import (
	"fmt"
	"math/rand"
	"crypto/sha1"
	"encoding/hex"
)

type Kademlia struct {
	Rtable *RoutingTable
	data []dataStruct
	net	Network
	hashmap map[string][]byte
}
type dataStruct struct {
	Hash	string
	Value	[]byte
}

func InitKad(me Contact) *Kademlia {
	node := &Kademlia{
		Rtable: NewRoutingTable(me),
		hashmap: make(map[string][]byte),
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
	doublet := make(map[Contact]bool)
	for _, mine := range myClosest {
		shortlist = append(shortlist, mine)
		doublet[mine] = true
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
		fmt.Printf("\nrunningRoutines: %d\n", runningRoutines)
		for _, candidate := range recived {
			if doublet[candidate] != true && !(candidate.Address == network.Contact.Address) && !(candidate.ID == nil) {
				doublet[candidate] = true
				candidate.CalcDistance(network.Contact.ID)
				shortlist = append(shortlist, candidate)
			}
		}
		shortlist = qsort(shortlist)
		runningRoutines--
		if closestNode != shortlist[0] {
			closestNode = shortlist[0]
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
	shortlist = qsort(shortlist)
	if len(shortlist) > 20 {
		result <- shortlist[:20]
	} else {
		result <- shortlist
	}
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO iterativeFindValue

}


func (kademlia *Kademlia) Store(data []byte, network *Network) {
	ch := make(chan []Contact)
	contact := NewContact(NewKademliaID(HashData(data)), "")
	fmt.Println(contact.ID)
	go kademlia.LookupContact(*network, ch, contact)
	done :=  <- ch
	for _, c := range done{
		kademlia.net.SendStoreMessage(&c, data)
	}
}

func HashData(data []byte) string{
	hashedData := sha1.Sum(data)
	hashedStringdata := hex.EncodeToString(hashedData[0:])
	return hashedStringdata
}

func qsort(contact []Contact) []Contact {
	if len(contact) < 2 {
        return contact
    }
      
    left, right := 0, len(contact)-1
      
    pivot := rand.Int() % len(contact)
      
    contact[pivot], contact[right] = contact[right], contact[pivot]
      
    for i := range contact {
        if (contact[i].distance.Less(contact[right].distance)) {
            contact[left], contact[i] = contact[i], contact[left]
            left++
        }
    }
      
    contact[left], contact[right] = contact[right], contact[left]
      
    qsort(contact[:left])
    qsort(contact[left+1:])
      
    return contact
}

