package main

import (
	d "D7024E"
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	//Read port from args
	myport := "1"
	if len(os.Args) > 1 {
		myport = os.Args[1]
	}
	iPort, err := strconv.Atoi(myport)
	ErrorHandler(err)
	myip, err := externalIP()
	ErrorHandler(err)
	if myip == "" {
		myip = "127.0.0.1"
	}
	me := d.NewContact(d.NewRandomKademliaID(), myip+":"+myport)
	fmt.Println("I am: ", me.ID, me.Address+"\n")

	newNode := d.InitNode(myip, iPort, &me)
	//Read ip & node from args (node to join)
	bIP := ""
	//bNode := ""
	//Check if a known bootstrap node (c) was given
	if len(os.Args[1:]) == 2 {
		bIP = os.Args[2]
		//bNode = os.Args[3] we dont need to know nodeID only ip:port
		if bIP != "" {
			//Make contact of bootstrap node.
			bContact := d.NewContact(nil, bIP)

			//RPC PING node c and update buckets
			newNode.SendPingMessage(&bContact)

			//Check if my bucket was updated
			myContacts := newNode.Kad.Rtable.FindClosestContacts(d.NewKademliaID("0000000000000000000000000000000000000000"), 160)
			if len(myContacts) == 0 {
				ErrorHandler(errors.New("pinging bootstrap failed or buckets weren't updated"))
			}

			//iterativeFindNode for new node n
			val := newNode.IterativeFindNode()
			for _, c := range val {
				newNode.Kad.Rtable.AddContact(c)
			}


			//Update the k-buckets further away than the one bootstrap node falls in
			/*
				for i := 160; i > newNode.Kad.Rtable.GetBucketIndex(findBootstrap[0].ID); i-- {
					for k := 20; k < 20; k++ {
						newNode.SendPingMessage(&findBootstrap[i])

					}
				}*/
		}

	}
	out := make(chan []d.Contact)
	wg.Add(2)
	go newNode.Listen(me, iPort) //Handle any RPC
	go func() {                  //Handle cli at the same time as RCP
		cli := bufio.NewScanner(os.Stdin)
		fmt.Println("Listening on cli")
		for {
			cli.Scan()
			text := cli.Text()
			if len(text) != 0 {
				fmt.Println(text)
				if text[:4] == "PING" {
					node := d.NewContact(nil, text[5:])
					newNode.SendPingMessage(&node)
				}
				if text[:4] == "FIND" {
					node := d.NewContact(d.NewKademliaID(text[5:]), text[46:])
					go newNode.SendFindContactMessage(&node, out)
					x := <-out
					fmt.Println("Got following contacts: ")
					fmt.Println(x)
				}
				/*Placeholder
				PUT
				GET
				EXIT
				 */
				if text[:3] == "PUT"{
					storeData := []byte(text[4:])
					fmt.Println("Storing data on other nodes")
					newNode.Kad.Store(storeData)

				}
				if text[:3] == "GET"{
					fmt.Println("Fetching data...")
					newNode.SendFindDataMessage("")
				}
				if text[:4] == "EXIT" {
					fmt.Println("Node is shutting down in 3 seconds...")
					time.Sleep(3 * time.Second)
					os.Exit(0)
				}

			}
		}
	}()
	wg.Wait()

}
//ErrorHandler to not fill code with if statements
func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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
			if strings.Index(ip.String(), "10.") == 0 { //IP should start at 10. within containers
				return ip.String(), nil
			}

		}
	}
	return "", err
}
