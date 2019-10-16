package D7024E

import (
	"testing"
	"time"
)

func TestKademliaID_Less(t *testing.T) {
	testID1 := NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	testID2 := NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE")
	//Should give value "false"
	falseCase := (testID1.Less(testID2))
	if falseCase != false {
		t.Errorf("testID1.Less(testID2) = %v; want false", falseCase)
	}

}
func TestKademliaID_Equals(t *testing.T) {
	testID1 := NewRandomKademliaID()
	time.Sleep(50 * time.Millisecond)
	testID2 := NewRandomKademliaID()
	//Should give value "false"
	falseCase := (testID1.Equals(testID2))
	if falseCase == true {
		t.Errorf("testID1.Equals(testID2) = %v; want false", falseCase)
	}
}
func TestKademliaID_CalcDistance(t *testing.T) {
	testID1 := NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	testID2 := NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE")
	//Expected result of this is 00.......1
	testXor := testID1.CalcDistance(testID2)
	if testXor.String() != "0000000000000000000000000000000000000001" {
		t.Errorf("testID1.CalcDistance(testID2) = %v; want 0000000000000000000000000000000000000001", testXor)
	}
}
