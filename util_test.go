package D7024E

import (
	"fmt"
	"reflect"
	"testing"
)

func TestByteToContact(t *testing.T) {
	c := make([]Contact, 0)
	c = append(c, NewContact(NewKademliaID("ffffffff00000000000000000000000000000000"), "localhost:1"))
	fmt.Println(c[0].ID.String())
	tests := []struct {
		name string
		args []byte
		want []Contact
	}{
		{
			"ByteToContact1", []byte(c[0].ID.String() + " " + c[0].Address + " "), c,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToContact(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteToContact() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestContactToByte(t *testing.T) {
	newCont := Contact{
		distance: NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
		ID:NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
		Address:"1",
	}
	newWhat := []byte(newCont.ID.String() + " "+ newCont.Address + " " + newCont.distance.String()+"\n")
	newArr := make([]byte,0)
	newArr = append(newArr,newWhat[:]... )
	contactarr := []Contact{}
	contactarr = append(contactarr, newCont)
	fmt.Println(contactarr[0])
	fmt.Println(newArr)
	tests := []struct {
		args []Contact
		want []byte
	}{
		{contactarr,newArr},
	}
	for _, tt := range tests {
		t.Run(tt.args[0].Address, func(t *testing.T) {
			if got := ContactToByte(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContactToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashData(t *testing.T) {
	fmt.Println("Testing the hasher: ")
	testArray := []byte("123")
	svar := HashData(testArray)
	if svar == "40bd001563085fc35165329ea1ff5c5ecbdbbeef" {
		fmt.Println("Correct Hash for " + string(testArray) + "!")
		fmt.Println(svar)
	} else {
		fmt.Println("Not hashed correctly!")
	}
}

func TestQuicksort(t *testing.T) {

	target := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"), "localhost:8000")
	node1 := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFA"), "localhost:8001")
	node2 := NewContact(NewKademliaID("EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE"), "localhost:8002")
	node3 := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE"), "localhost:8003")
	node4 := NewContact(NewKademliaID("1111111111111111111111111111111111111111"), "localhost:8004")
	var unsortedList []Contact
	var sortedList []Contact
	unsortedList = append(unsortedList, node1, node2, node3, node4)
	sortedList = append(sortedList, node3, node1, node2, node4)

	quickSortedList := qsort(unsortedList, target)

	for i := range quickSortedList {
		t.Run("sorting test", func(t *testing.T) {
			if got := quickSortedList[i].ID.String(); !reflect.DeepEqual(got, sortedList[i].ID.String()) {
				t.Errorf("quickSortedList = %v, want %v", got, sortedList[i])
			}
		})
	}
}
