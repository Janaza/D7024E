package D7024E

import "fmt"

type Kademlia struct {
	Rtable *RoutingTable
}

func InitKad(me Contact) *Kademlia {
	node := &Kademlia{
		Rtable: NewRoutingTable(me),
	}
	return node
}

func (kademlia *Kademlia) LookupContact(net Network, result chan []Contact, target Contact) {

	// TODO iterativeFindNode
	alpha := 3
	found := make(chan []Contact)

	//shortlist := make([]Contact, 0)
	var closestNode Contact
	myClosest := net.Kad.Rtable.FindClosestContacts(target.ID, alpha)
	closestNode = myClosest[0]
	shortlist := ContactCandidates{}
	doublet := make(map[Contact]bool)
	for _, mine := range myClosest {
		shortlist.contacts = append(shortlist.contacts, mine)
		doublet[mine] = true
	}
	runningRoutines := 0
	for runningRoutines < 3 && len(shortlist.contacts) > 1 {
		runningRoutines++
		net.SendFindContactMessage(&shortlist.contacts[runningRoutines], found)
	}
	if len(shortlist.contacts) == 1 {
		net.SendFindContactMessage(&shortlist.contacts[runningRoutines], found)
		runningRoutines++
	}

	for runningRoutines > 0 {
		recived := <-found

		for _, candidate := range recived {
			if doublet[candidate] != true {
				doublet[candidate] = true
				fmt.Println(candidate.String())
				shortlist.contacts = append(shortlist.contacts, candidate)
			}
		}

		shortlist.Sort()
		runningRoutines--
		if closestNode != shortlist.contacts[0] {
			closestNode = shortlist.contacts[0]
			for i := range shortlist.contacts[:2] {
				runningRoutines++
				fmt.Println(shortlist.contacts[i])
				net.SendFindContactMessage(&shortlist.contacts[i], found)

			}
		}
	}
	if len(shortlist.contacts) > 20 {
		for _, c := range shortlist.contacts {
			fmt.Println("Result: " + c.String())
		}
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
