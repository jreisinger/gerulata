package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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

func getJSON(nodesByID map[int]*Node) ([]byte, error) {
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

	return json.Marshal(&nodes)
}

func getNodesByID(r io.Reader) (map[int]*Node, error) {
	nodesByID := make(map[int]*Node)

	var id int
	var n int
	var err error

	input := bufio.NewScanner(r)

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
				return nil, fmt.Errorf("'%s' is not an ID (%v)", line, err)
			}
			_, ok := nodesByID[id]
			if ok {
				return nil, fmt.Errorf("duplicate ID: %d", id)
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

	return nodesByID, input.Err()
}
