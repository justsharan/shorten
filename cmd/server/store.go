package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readStore reads the store file and collects all name->URL mappings
func readStore(file string) (map[string]string, error) {
	osFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer osFile.Close()

	routes := map[string]string{}

	sc := bufio.NewScanner(osFile)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		res := strings.Split(sc.Text(), " ")
		if len(res) > 1 {
			routes[res[0]] = res[1]
		}
	}

	return routes, nil
}

// writeStore writes routes to the store file
func writeStore(file string, routes map[string]string) error {
	osFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer osFile.Close()

	for key, val := range routes {
		osFile.WriteString(fmt.Sprintf("%s %s\n", key, val))
	}

	return nil
}
