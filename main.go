package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nodesByID, err := getNodesByID(file)
	if err != nil {
		log.Fatal(err)
	}

	j, err := getJSON(nodesByID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", j)
}
