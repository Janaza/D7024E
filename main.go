package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("hello world")
	Listener()
}
func ErrorHandler(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func Listener(){
	fmt.Println("Testing Listen...")
	servr, err := net.ListenUDP("udp", &net.UDPAddr{IP:[]byte{0,0,0,0},Port:10001,Zone:""})
	ErrorHandler(err)
	//defer servr.Close()
	for{
		msgbuf := make([]byte, 1024)
		n, adrs, err := servr.ReadFrom(msgbuf)
		ErrorHandler(err)
		fmt.Println("Msg from a friend: ", string(msgbuf[0:n])," from ", adrs)
	}



}
