package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/jreisinger/checkip/checks"
)

type Node struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Type      string   `json:"type"`
	Activity  string   `json:"activity"`
	Impact    string   `json:"impact"`
	Influence string   `json:"influence"`
	Threat    string   `json:"threat"`
	UrlPath   string   `json:"url"`
	IPs       []net.IP `json:"ip_addresses"`
	AS        string   `json:"as"`
}

// var as = flag.Bool("as", false, "get AS description")

func main() {
	flag.Parse()

	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nodesByID, err := getNodesByID(file)
	if err != nil {
		log.Fatal(err)
	}

	// if *as {
	// 	a, err := getAS()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(a)

	j, err := getJSON(nodesByID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", j)

}

func extractUrl(s string) (*url.URL, error) {
	r := regexp.MustCompile(`\(([^)]+)\)`)
	rawURL := r.FindStringSubmatch(s)[1]
	return url.Parse(rawURL)
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
			UrlPath:   node.UrlPath,
			IPs:       node.IPs,
			AS:        node.AS,
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
			if nodesByID[id].Type == "Web" {
				// Enrich with URL path
				u, err := extractUrl(nodesByID[id].Title)
				if err != nil {
					return nil, err
				}
				nodesByID[id].UrlPath = u.Path

				// Enrich with IP addresses
				ips, _ := net.LookupIP(nodesByID[id].UrlPath) // NOTE: ignoring lookup error
				nodesByID[id].IPs = ips

				// Enrich with AS
				if len(ips) != 0 {
					as, err := getAS(ips[0])
					if err != nil {
						return nil, err
					}
					nodesByID[id].AS = as
				}
			}
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

func getAS(ip net.IP) (string, error) {
	res, err := checks.CheckAS(ip)
	if err != nil {
		return "", err
	}
	j, err := res.Info.JsonString()
	if err != nil {
		return "", err
	}
	sr := strings.NewReader(j)
	var a checks.AS
	if err := json.NewDecoder(sr).Decode(&a); err != nil {
		return "", err
	}
	return a.Description, nil
}
