package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Node struct {
	Title     string `json:"title"`
	Type      string `json:"type"`
	Activity  string `json:"activity"`
	Impact    string `json:"impact"`
	Influence string `json:"influence"`
	Threat    string `json:"threat"`
}

type Nodes map[int]*Node // map of node ID into Node

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nodes := make(map[int]*Node)

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
			_, ok := nodes[id]
			if ok {
				log.Fatalf("duplicate ID: %d", id)
			}
			nodes[id] = &Node{}
		case n == 2:
			nodes[id].Title = line
		case n == 3:
			nodes[id].Type = line
		case n == 4:
			nodes[id].Activity = line
		case n == 5:
			nodes[id].Impact = line
		case n == 6:
			nodes[id].Influence = line
		case n == 7:
			nodes[id].Threat = line
			n = 0
		}
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(&nodes); err != nil {
		log.Fatal(err)
	}
}
