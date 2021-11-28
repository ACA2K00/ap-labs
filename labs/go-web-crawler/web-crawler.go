package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch5/links"
)


var tokens = make(chan struct{}, 20)
var depth = 0

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} 
	list, err := links.Extract(url)
	<-tokens 

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {

	if len(os.Args) < 3 {
		println("Incorrect Number of Arguments: ./web-crawler -depth=n url")
		os.Exit(1)
	}
	lista := []string{}
	d := strings.Split(os.Args[1], "=")
	if d[0] != os.Args[1] {
		dtmp, flag := strconv.Atoi(d[1])
		if d[0] == "-depth" && flag == nil {
			depth = dtmp
			lista = append(lista, d[1])
		} else {
			println("Incorrect Parameters: ./web-crawler -depth=n url")
			os.Exit(2)
		}
	} else {
		println("Incorrect Parameters: ./web-crawler -depth=n url")
		os.Exit(3)
	}
	lista = append(lista, os.Args[2:]...)
	worklist := make(chan []string)
	var n int 

	// Start with the command-line arguments.
	n++

	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		depthOnWorklist := 0
		for num, link := range list {
			if num == 0 {
				depthOnWorklist, _ = strconv.Atoi(link)
				if depthOnWorklist <= depth {
					continue
				} else {
					break
				}
			}
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					linkurl := crawl(link)
					linkurl = append([]string{strconv.Itoa(depthOnWorklist + 1)}, linkurl...)
					worklist <- linkurl
				}(link)
			}
		}
	}
	println("Task Completed Successfully")
}