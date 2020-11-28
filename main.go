package main

import (

    "fmt"
	"log"

    "./network"

    //"./configuration"
)

func main() {
    connectedInterfaces, err := network.GetConnectedWifiInterfaces()
    if err != nil {
        log.Fatal(err)
    
}


    
    fmt.Println(connectedInterfaces)
}


func panicError(err error){
    if err != nil {
        log.Panic(err)
    }
}

func fatalError(err error){
    if err != nil {
        log.Fatal(err)
    }
}

