package main

 import (
         "log"
         "net"
 )

 func main() {
         host := "localhost"
         port := "10001"

         service := host + ":" + port

         RemoteAddress, err := net.ResolveUDPAddr("udp", service)

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
