package D7024E

import "strings"

func ByteToContact(msg []byte) []Contact {
	s := string(msg)
	slice := strings.Split(s, "\n")
	arr := make([]Contact, 20)
	//var contact Contact
	for i, line := range slice {
		if len(line) != 0 {
			contact := NewContact(NewKademliaID(line[:40]), line[41:41+strings.Index(line[41:], " ")])
			if len(line[41+strings.Index(line[41:], " "):]) > 2 {
				contact.distance = NewKademliaID(line[41+strings.Index(line[41:], " "):])
			}
			arr[i] = contact
		}
	}
	return arr
}

func ContactToByte(contactArr []Contact) []byte {
	closeCToByte := make([]byte, 0)
	closeContactsByte := make([]byte, 0)
	for i := 0; i < len(contactArr); i++ {
		closeCToByte = []byte(contactArr[i].ID.String() + " " + contactArr[i].Address + " " + contactArr[i].distance.String() + "\n")
		closeContactsByte = append(closeContactsByte, closeCToByte[:]...)
	}

	return closeContactsByte

}
