package main

import (
	"encoding/json"
	"net"
	"net/url"
	"regexp"
	"strings"

	"github.com/jreisinger/checkip/check"
)

func enrich(node *Node) error {
	// Enrich with URL path
	u, err := extractUrl(node.Title)
	if err != nil {
		return err
	}
	node.UrlPath = u.Path

	// Enrich with IP addresses
	ips, _ := net.LookupIP(node.UrlPath) // NOTE: ignoring lookup error
	node.IPs = ips

	if len(ips) != 0 {
		// Enrich with IP address availability
		p, err := ping(ips[0])
		if err != nil {
			return err
		}
		node.Ping = p

		// Enrich with AS
		as, err := getAS(ips[0])
		if err != nil {
			return err
		}
		node.AS = as
	}

	return nil
}

func ping(ip net.IP) (string, error) {
	res, err := check.Ping(ip)
	if err != nil {
		return "", err
	}
	return res.Info.Summary(), nil
}

func getAS(ip net.IP) (string, error) {
	res, err := check.IPtoASN(ip)
	if err != nil {
		return "", err
	}
	j, err := res.Info.Json()
	if err != nil {
		return "", err
	}
	sr := strings.NewReader(string(j))
	var a check.AutonomousSystem
	if err := json.NewDecoder(sr).Decode(&a); err != nil {
		return "", err
	}
	return a.Description, nil
}

func extractUrl(s string) (*url.URL, error) {
	r := regexp.MustCompile(`\(([^)]+)\)`)
	rawURL := r.FindStringSubmatch(s)[1]
	return url.Parse(rawURL)
}
