package D7024E

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
)

type Kademlia struct {
	Rtable *RoutingTable
	net     Network
	hashmap map[string][]byte
}

type Shortlist struct {
	ls []Contact
	v  map[string]bool
}

func InitKad(me Contact) *Kademlia {
	node := &Kademlia{
		Rtable:  NewRoutingTable(me),
		hashmap: make(map[string][]byte),
	}
	return node
}

const a = 3

func (kademlia *Kademlia) LookupContact(network Network, result chan []Contact, target Contact) {
	var closestNode Contact
	var x []Contact
	found := make(chan []Contact)
	doublet := make(map[string]bool)
	iRoutines := 0
	myClosest := network.Kad.Rtable.FindClosestContacts(target.ID, a)
	closestNode = myClosest[0]
	sl := &Shortlist{
		ls: make([]Contact, 0),
		v:  make(map[string]bool),
	}

	for _, mine := range myClosest {
		sl.insert(false, mine)
		doublet[mine.ID.String()] = true
	}

	for iRoutines < a && iRoutines < len(sl.ls) {
		go network.SendFindContactMessage(&sl.ls[iRoutines], found, sl)
		x = <-found
		sl.v[sl.ls[iRoutines].ID.String()] = true
		iRoutines++
	}

	for iRoutines > 0 {
		recived := x
		for _, candidate := range recived {
			if !(candidate.Address == network.Contact.Address) && !(candidate.ID == nil) {
				if doublet[candidate.ID.String()] == false {
					doublet[candidate.ID.String()] = true
					candidate.CalcDistance(network.Contact.ID)
					sl.insert(false, candidate)
				}
			}
		}
		sl.ls = qsort(sl.ls, target)
		iRoutines--
		if closestNode.ID.String() != sl.ls[0].ID.String() {
			closestNode = sl.ls[0]
			for i := range sl.ls {
				if i >= a || i >= len(sl.ls) {
					break
				}
				if sl.v[sl.ls[i].ID.String()] == false {
					iRoutines++
					go network.SendFindContactMessage(&sl.ls[i], found, sl)
					x = <-found
					sl.v[sl.ls[i].ID.String()] = true
				}

			}
		}
	}
	for i, c := range sl.ls {
		if i >= len(sl.ls) {
			break
		}
		if sl.v[sl.ls[i].ID.String()] == false {
			go network.SendFindContactMessage(&c, found, sl)
			x = <-found
			sl.v[c.ID.String()] = true
		}
	}
	sl.ls = qsort(sl.ls, target)
	if len(sl.ls) > 20 {
		result <- sl.ls[:20]
	} else {
		result <- sl.ls
	}
}

func (kademlia *Kademlia) LookupData(network Network, target Contact, hash string) string {
	alpha := 3
	value := make(chan string)
	found := make(chan []Contact)
	var x []Contact
	var y string
	myClosest := network.Kad.Rtable.FindClosestContacts(target.ID, alpha)
	var shortlist []Contact
	var noKeyShortlist []Contact
	doublet := make(map[string]bool)
	visited := make(map[string]bool)
	for _, mine := range myClosest {
		shortlist = append(shortlist, mine)
		doublet[mine.ID.String()] = true
	}

	runningRoutines := 0
	for runningRoutines < 3 && len(shortlist) > 1 {
		go network.SendFindDataMessage(hash, &shortlist[runningRoutines], found, value)
		x = <-found
		y = <-value
		if y != "" {
			if len(noKeyShortlist) > 0 {
				fmt.Println("Storing at closest contact")
				kademlia.net.SendStoreMessage(&noKeyShortlist[0], []byte(y))
			}
			runningRoutines = 0
			return y
		}
		runningRoutines++
		for _, i := range x {
			if i.ID != nil {
				i.CalcDistance(target.ID)
				noKeyShortlist = append(noKeyShortlist, i)
			}
		}
		noKeyShortlist = qsort(noKeyShortlist, target)
	}

	if len(shortlist) == 1 {
		runningRoutines++
		go network.SendFindDataMessage(hash, &shortlist[0], found, value)
		x = <-found
		y = <-value
		if y != "" {
			if len(noKeyShortlist) > 0 {
				fmt.Println("Storing at closest contact")
				kademlia.net.SendStoreMessage(&noKeyShortlist[0], []byte(y))
			}
			runningRoutines = 0
			return y
		}
		for _, i := range x {
			if i.ID != nil {
				i.CalcDistance(target.ID)
				noKeyShortlist = append(noKeyShortlist, i)
			}
		}
		noKeyShortlist = qsort(noKeyShortlist, target)

	}

	for runningRoutines > 0 && len(x) > 0 {
		recived := x
		for _, candidate := range recived {
			if !(candidate.Address == network.Contact.Address) && !(candidate.ID == nil) {
				if doublet[candidate.ID.String()] == false {
					doublet[candidate.ID.String()] = true
					candidate.CalcDistance(target.ID)
					shortlist = append(shortlist, candidate)

				}
			}
		}
		shortlist = qsort(shortlist, target)
		runningRoutines--
		for i := range shortlist {

			if visited[shortlist[i].ID.String()] == false {
				visited[shortlist[i].ID.String()] = true
				runningRoutines++
				go network.SendFindDataMessage(hash, &shortlist[i], found, value)
				x = <-found
				y = <-value
				if y != "" {
					if len(noKeyShortlist) > 0 {
						fmt.Println("Storing at closest contact")
						kademlia.net.SendStoreMessage(&noKeyShortlist[0], []byte(y))
					}
					runningRoutines = 0
					return y
				}
				for _, i := range x {
					if i.ID != nil {
						i.CalcDistance(target.ID)
						noKeyShortlist = append(noKeyShortlist, i)
					}
				}
				noKeyShortlist = qsort(noKeyShortlist, target)
			}
		}
	}

	shortlist = qsort(shortlist, target)

	var shortlistString string

	if len(shortlist) > 20 {
		for _, i := range shortlist[:20] {
			shortlistString = shortlistString + i.String() + "\n"
		}
	} else {
		for _, i := range shortlist {
			shortlistString = shortlistString + i.String() + "\n"
		}
	}

	return shortlistString
}

func (kademlia *Kademlia) Store(data []byte, network *Network) {
	ch := make(chan []Contact)
	contact := NewContact(NewKademliaID(HashData(data)), "")
	//fmt.Println(contact.ID)
	go kademlia.LookupContact(*network, ch, contact)
	done := <-ch
	for _, c := range done {
		kademlia.net.SendStoreMessage(&c, data)
	}
}

func HashData(data []byte) string {
	hashedData := sha1.Sum(data)
	hashedStringdata := hex.EncodeToString(hashedData[0:])
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

func (sl *Shortlist) insert(v bool, c Contact) []Contact {
	sl.ls = append(sl.ls, c)
	return sl.ls
}

func (sl *Shortlist) removeContact(c Contact) {
	for i, f := range sl.ls {
		if f.ID.String() == c.ID.String() {
			copy(sl.ls[i:], sl.ls[i+1:])
			sl.ls = sl.ls[:len(sl.ls)-1]
			return
		}

	}
	fmt.Println("contact not in list")
}
