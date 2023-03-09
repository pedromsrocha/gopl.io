package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	scanners := make([]*bufio.Scanner, len(os.Args[1:]))
	for i, v := range os.Args[1:] {
		conn, err := net.Dial("tcp", v)
		if err != nil {
			log.Fatal(err)
		}
		scanners[i] = bufio.NewScanner(conn)
	}

	for {
		times := ""
		for _, s := range scanners {
			if !s.Scan() {
				return
			}
			times += s.Text() + " "
		}
		fmt.Printf("%s\n", times)
	}
}
