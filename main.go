package main

import (
  "flag"
  "os"
  "./lib"
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
        lib.RunHost(connIP)
    } else {
        connIP := os.Args[1]
        lib.RunGuest(connIP)        
    }
}

