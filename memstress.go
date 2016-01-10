package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var a = []string{}

func main() {
	b, _ := ioutil.ReadFile("/etc/hosts")
	s := string(b)

	for {
		a = append(a, fmt.Sprintf("%s-%s", s, time.Now().UTC()))
	}
}
