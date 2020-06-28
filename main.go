package main

import (
	"fmt"
	"log"

	"github.com/henrikkorsgaard/wifi"
)

func main() {

	fmt.Println("hello")
	c, err := wifi.New()
	defer c.Close()

	if err != nil {
		log.Fatal(err)
	}

	phys, err := c.PHYs()
	if err != nil {
		log.Fatal(err)
	}

	for _, phyInterface := range phys {
		fmt.Println(phyInterface)
	}
}
