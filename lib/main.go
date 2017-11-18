package lib

import (
  "fmt"
  "net"
  "log"
  "bufio"
)

const port = "8080"

// RunHost takes an argument of ip address string and listens on that address
func RunHost(ip string) {
    ipAndPort := ip + ":" + port
    listener, listenErr := net.Listen("tcp", ipAndPort)
    if listenErr != nil {
        log.Fatal("Error: ", listenErr)
    }

    fmt.Println("Listening on ", ipAndPort)

    conn, acceptErr := listener.Accept()
    if acceptErr != nil {
        log.Fatal("Error: ", acceptErr)
    }

    fmt.Println("New connection accepted")

    reader := bufio.NewReader(conn)
    message, readErr := reader.ReadString('\n')
    if readErr != nil {
      log.Fatal("Error: ", readErr)
    }

    fmt.Println("message received: ", message)
}

// RunGuest takes an argument of ip address string as the destination to send message
func RunGuest(ip string) {

}