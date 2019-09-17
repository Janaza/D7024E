package d7024e

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func Udpclient() {
	remoteAddr := bufio.NewScanner(os.Stdin)
	arr := make([]string, 0)
	for {
		fmt.Print("Enter IP: ")
		remoteAddr.Scan()
		text := remoteAddr.Text()
		if len(text) != 0 {
			fmt.Println(text)
			arr = append(arr, text)
		} else {
			break
		}
	}
	RemoteAddress, err := net.ResolveUDPAddr("udp", strings.Join(arr, ""))

	connection, err := net.DialUDP("udp", nil, RemoteAddress)
	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()

	message := []byte("Hello server =)")
	_, err = connection.Write(message)

	if err != nil {
		log.Println(err)
	}

}
