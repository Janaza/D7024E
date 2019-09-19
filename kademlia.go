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

func (kademlia *Kademlia) LookupContact(net Network, result chan []Contact, target Contact) {
	// TODO iterativeFindNode
	alpha := 3
	found := make(chan Contact)

	shortlist := make([]Contact, 0)
	var closestNode Contact
	myClosest := net.Kad.Rtable.FindClosestContacts(target.ID, alpha)
	doublet := make(map[Contact]bool)
	for _, mine := range myClosest {
		shortlist = append(shortlist, mine)
		doublet[mine] = true
	}
	runningRoutines := 0
	for runningRoutines < 3 {
		runningRoutines++
		go net.SendFindContactMessage(&shortlist[runningRoutines], found) //add channel
	}

	for runningRoutines > 0 {
		recived := <-found
		for _, candidate := range newCandidates {
			if d[candidate] != true {
				shortlist = append(shortlist, candidate)
			}
		}

	}

}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO iterativeFindValue

}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO iterativeStore
}
