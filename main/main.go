package main

import (
	"bufio"
	"d7024e"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	//Read port from args
	myport := os.Args[1]
	iPort, err := strconv.Atoi(myport)
	if err != nil {
		log.Fatal(err)
	}
	myip, err := externalIP()
	if err != nil {
		log.Fatal(err)
	}
	me := d7024e.NewContact(d7024e.NewRandomKademliaID(), myip+":"+myport)
	fmt.Println("I am: ", me.ID, me.Address)
	newNode := d7024e.InitNode(me, iPort)

	//Set newNode up for RCP listeing

	//Read ip & node from args (node to join)
	bIP := ""
	bNode := ""
	if len(os.Args[1:]) == 3 {
		//A known bootstrap node (c) was given
		bIP = os.Args[2]
		bNode = os.Args[3]
		if bIP != "" && bNode != "" {
			//Make contact of bootstrap node.
			bContact := d7024e.NewContact(d7024e.NewKademliaID(bNode), bIP)
			//RPC PING node c and update buckets
			//newNode.Network.SendPingMessage(&bContact)

			//Example
			newNode.Network.SendPing(&bContact)
			//Example

			//iterativeFindNode for new node n
			//newNode.LookupContact(&me)
		}

	}
	wg.Add(2)
	go d7024e.Listen("127.0.0.1", iPort) //Handle any RPC change ip to myip (externalIP())
	go func() {                          //Handle cli at the same time as RCP
		cli := bufio.NewScanner(os.Stdin)
		for {
			fmt.Println("cli: ")
			cli.Scan()
			text := cli.Text()
			if len(text) != 0 {
				fmt.Println(text)
				if text[:4] == "PING" {
					node := (d7024e.NewContact(nil, text[5:]))
					newNode.Network.SendPing(&node)
				}
			}
		}
	}()
	wg.Wait()

}
func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//get my outbound ip
/*func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
*/
func externalIP() (string, error) {
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
			if strings.Index(ip.String(), "10.") == 0 { //IP should start at 10.
				return ip.String(), nil
			}

		}
	}
	return "", err
}

/*func Listener() {
	fmt.Println("Server is starting...")
	hej, _ := externalIP()
	fmt.Println(hej)
	//myip := GetOutboundIP()
	//ip must be of type string
	//myip2 := myip.String()
	listenAdrs, _ := net.ResolveUDPAddr("udp", hej+":10001")
	servr, err := net.ListenUDP("udp", listenAdrs)
	ErrorHandler(err)
	defer servr.Close()
	for {
		fmt.Println(net.Interfaces())
		fmt.Println("Listening on: " + string(listenAdrs.String()))
		msgbuf := make([]byte, 1024)
		n, adrs, err := servr.ReadFrom(msgbuf)
		ErrorHandler(err)
		fmt.Println("Msg from a friend: ", string(msgbuf[0:n]), " from ", adrs)
	}

}*/
