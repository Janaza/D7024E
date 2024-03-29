package D7024E

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"net"
	"strings"
)

func Eth0IP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		if iface.Name == "eth0" {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				if ip == nil || ip.IsLoopback() {
					continue
				}
				ip = ip.To4()
				if ip == nil {
					continue // not an ipv4 address
				}
				return ip.String(), nil

			}
		}
	}
	return "localhost", err
}

func ByteToContact(msg []byte) []Contact {
	s := string(msg)
	slice := strings.Split(s, "\n")
	arr := make([]Contact, 0)
	//var contact Contact
	for _, line := range slice {
		if len(line) != 0 {
			contact := NewContact(NewKademliaID(line[:40]), line[41:41+strings.Index(line[41:], " ")])
			if len(line[41+strings.Index(line[41:], " "):]) > 2 {
				contact.distance = NewKademliaID(line[41+strings.Index(line[41:], " "):])
			}
			arr = append(arr, contact)
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
