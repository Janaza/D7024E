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

func Listener(){
	fmt.Println("Server is starting...")
	listenAdrs, _ := net.ResolveUDPAddr("udp", "localhost:10001")
	servr, err := net.ListenUDP("udp", listenAdrs)
	ErrorHandler(err)
	defer servr.Close()
	for{
		fmt.Println("Listening on: " + string(listenAdrs.String()))
		msgbuf := make([]byte, 1024)
		n, adrs, err := servr.ReadFrom(msgbuf)
		ErrorHandler(err)
		fmt.Println("Msg from a friend: ", string(msgbuf[0:n])," from ", adrs)
	}



}
