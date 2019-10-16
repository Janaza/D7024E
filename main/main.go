package main

//powershell.exe -executionpolicy bypass .\run.ps1
import (
	d "D7024E"
	"flag"
	"os"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	myport := flag.String("port", "7070", "Provide any port, if unset defaluts to 7070")
	bootstrap := flag.String("bootstrapIP", "", "Provide the ip of any node to join network, if unset this node is asumed to be bootstrap")
	flag.Parse()
	var newNode *d.Network

	if *bootstrap != "" {
		newNode = d.InitJoin(*myport, *bootstrap)
	} else {
		newNode = d.InitBootstrap(*myport)
	}

	wg.Add(2)
	go newNode.Listen(*newNode.Contact)
	go newNode.CliHelper(os.Stdin)
	wg.Wait()

}
