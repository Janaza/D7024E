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
	candidates := make([]Contact, 20)
	myClosest := net.Kad.Rtable.FindClosestContacts(target.ID, alpha)
	doublet := make(map[Contact]bool)
	for _, mine := range myClosest {
		candidates = append(candidates, mine)
		doublet[mine] = true
	}
	runningRoutines := 0
	for runningRoutines < 3 {
		runningRoutines++
		go net.SendFindContactMessage(&candidates[runningRoutines]) //add channel
	}

	for runningRoutines > 0 {
		/*listen on go routine channel
		newCandidates := <- channel
		for _, candidate := range newCandidates {
			if d[candidate] != true {
			candidates = append(candidates, candidate)
			}
		}


		*/
	}

}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO iterativeFindValue

}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO iterativeStore
}
