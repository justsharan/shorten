package main

import (
	"io/ioutil"
	"strings"
)

func readFile(path string, routes *map[string]string) error {
	bytes, err := ioutil.ReadFile("routes.txt")
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(bytes), "\n") {
		res := strings.Split(line, " ")
		if len(res) > 1 {
			(*routes)[res[0]] = res[1]
		}
	}

	return nil
}

func saveFile(path string, routes *map[string]string) {
	var content string
	for k, v := range *routes {
		content += k + " " + v + "\n"
	}
	ioutil.WriteFile(path, []byte(content), 0466)
}
