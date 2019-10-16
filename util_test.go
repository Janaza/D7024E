package D7024E

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestEth0IP(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Eth0IP()
			if (err != nil) != tt.wantErr {
				t.Errorf("Eth0IP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Eth0IP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteToContact(t *testing.T) {
	type args struct {
		msg []byte
	}
	tests := []struct {
		name string
		args args
		want []Contact
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToContact(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
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
	var randomList []Contact
	target := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"), "localhost:8000")
	node1 := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFA"), "localhost:8001")
	node2 := NewContact(NewKademliaID("EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE"), "localhost:8002")
	node3 := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE"), "localhost:8003")
	node4 := NewContact(NewKademliaID("1111111111111111111111111111111111111111"), "localhost:8004")
	var contactList []Contact
	contactList = append(contactList, node1, node2, node3, node4)
	/*	//sortedList := qsort(contactList, target)
		fmt.Println("The closests contacts too " + target.ID.String() + " are: ")
		for i := range sortedList{
			fmt.Println(sortedList[i].ID.String())
		}*/

	for b := 0; b < 5; b++ {
		randomContact := NewContact(NewRandomKademliaID(), "localhost:800"+string(b))
		time.Sleep(1 * time.Second)
		randomList = append(randomList, randomContact)
	}
	fmt.Println("The randomized list: ")
	for a := range randomList {
		fmt.Println(randomList[a].ID)
	}
	sortedRandom := qsort(randomList, target)
	fmt.Println("Sorting target is fffff....")
	fmt.Println("The sorted list: ")
	for a := range sortedRandom {
		fmt.Println(sortedRandom[a].ID)
	}
}
