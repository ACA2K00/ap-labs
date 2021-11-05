package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func handleClientConn(address string) {

	separated := strings.Split(address, "=")

	conn, err := net.Dial("tcp", separated[1])

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan int)

	go func() {
		s := bufio.NewScanner(conn)
		for s.Scan() {
			fmt.Println(separated[0]+"\t"+":", s.Text())
		}
		log.Println("done")
		done <- 2
	}()

	x := 1
	x = <-done
	log.Println("Channel Closed with value: ", x)
	close(done)

}

func main() {
	ch := make(chan int)
	portsAndPlaces := os.Args[1:]
	for _, p := range portsAndPlaces {
		go handleClientConn(p)
	}

	for range portsAndPlaces {
		<-ch
	}

}