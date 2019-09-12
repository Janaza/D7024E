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
func Listener(){
	fmt.Println("Server is starting...")
	myip := GetOutboundIP()
	//ip must be of type string
	myip2 := myip.String()
	listenAdrs, _ := net.ResolveUDPAddr("udp", myip2+ ":10001")
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
