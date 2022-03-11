package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Node struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Activity  string `json:"activity"`
	Impact    string `json:"impact"`
	Influence string `json:"influence"`
	Threat    string `json:"threat"`
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nodesByID := make(map[int]*Node)

	var id int
	var n int
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		if line == "" {
			continue
		}

		n++
		switch {
		case n == 1: // new Node
			id, err = strconv.Atoi(line)
			if err != nil {
				log.Fatalf("'%s' is not an ID (%v)", line, err)
			}
			_, ok := nodesByID[id]
			if ok {
				log.Fatalf("duplicate ID: %d", id)
			}
			nodesByID[id] = &Node{}
		case n == 2:
			nodesByID[id].Title = line
		case n == 3:
			nodesByID[id].Type = line
		case n == 4:
			nodesByID[id].Activity = line
		case n == 5:
			nodesByID[id].Impact = line
		case n == 6:
			nodesByID[id].Influence = line
		case n == 7:
			nodesByID[id].Threat = line
			n = 0
		}
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}

	var nodes []Node

	for id, node := range nodesByID {
		n := Node{
			ID:        id,
			Title:     node.Title,
			Type:      node.Type,
			Activity:  node.Activity,
			Impact:    node.Impact,
			Influence: node.Influence,
			Threat:    node.Threat,
		}
		nodes = append(nodes, n)
	}

	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(&nodes); err != nil {
		log.Fatal(err)
	}
}
