package main

import (
  "fmt"
  "flag"
  "net"
  "os"
  "log"
  "bufio"
)

func main() {
    var isHost bool

    // The reason that we pass an identifier of isHost to BoolVar rather than just
    // the name of variable because we want to keep using the same memory address 
    // as the original var. This is because in the Parse() function, the var isHost is actually
    // assigned a value, should we remove the reference percent, then what would happen is
    // the var assigned value by the Parse() method is actually a copy of the isHost var declared above
    flag.BoolVar(&isHost, "listen", false, "Listens on the specified ip address")
    flag.Parse()

    if isHost {
        connIP := os.Args[2]
        runHost(connIP)
        // fmt.Println("is host")
    } else {
        fmt.Println("is guest")
    }
}

const port = "8080"

func runHost(ip string) {
    ipAndPort := ip + ":" + port
    listener, listenErr := net.Listen("tcp", ipAndPort)
    if listenErr != nil {
        log.Fatal("Error: ", listenErr)
    }

    conn, acceptErr := listener.Accept()
    if acceptErr != nil {
        log.Fatal("Error: ", acceptErr)
    }

    reader := bufio.NewReader(conn)
    message, readErr := reader.ReadString('\n')
    if readErr != nil {
      log.Fatal("Error: ", readErr)
    }

    fmt.Println("message received: ", message)
}

func runGuest(ip string) {

}