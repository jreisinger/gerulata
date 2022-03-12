package main

import (
	"fmt"
	"log"
	"os"
	"sync"
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

	// Enrich nodes of type Web with some useful data..
	var wg sync.WaitGroup
	for id := range nodesByID {
		if nodesByID[id].Type == "Web" {
			wg.Add(1)
			go func(node *Node) {
				defer wg.Done()
				if err := enrich(node); err != nil {
					log.Fatal(err)
				}
			}(nodesByID[id])
		}
	}
	wg.Wait()

	j, err := getJSON(nodesByID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", j)
}
