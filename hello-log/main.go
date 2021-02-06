package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "webserver.log", "SOme help on flag.")
	level := flag.String("verb", "GET", "HTTP verb to filter.")

	flag.Parse() // populates flag variables

	f, err := os.Open(*path) //"webserver.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // closes file on exit from function
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) { //"POST") {
			fmt.Print(s)
		}

	}
}
