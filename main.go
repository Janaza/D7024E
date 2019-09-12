package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	Listener()
}
func ErrorHandler(err error){
	if err != nil{
		log.Fatal(err)
	}
}
//get my outbound ip
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
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
			return ip.String(), nil
		}
	}
	return "", err
}

func Listener(){
	fmt.Println("Server is starting...")
	hej, _ := externalIP()
	fmt.Println(hej)
	//myip := GetOutboundIP()
	//ip must be of type string
	//myip2 := myip.String()
	listenAdrs, _ := net.ResolveUDPAddr("udp", hej+ ":10001")
	servr, err := net.ListenUDP("udp", listenAdrs)
	ErrorHandler(err)
	defer servr.Close()
	for{
		fmt.Println(net.Interfaces())
		fmt.Println("Listening on: " + string(listenAdrs.String()))
		msgbuf := make([]byte, 1024)
		n, adrs, err := servr.ReadFrom(msgbuf)
		ErrorHandler(err)
		fmt.Println("Msg from a friend: ", string(msgbuf[0:n])," from ", adrs)
	}



}
