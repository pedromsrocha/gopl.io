// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// !+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan struct {
		id string
		ch client
	})
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]string) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli.ch] = cli.id
			go announceCurrentClients(clients)

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

func announceCurrentClients(clients map[client]string) {
	messages <- "The following clients are in the chat:"
	for _, id := range clients {
		messages <- id
	}
}

// !+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	ch <- "Who are you?"
	input := bufio.NewScanner(conn)
	input.Scan()
	who := input.Text()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- struct {
		id string
		ch client
	}{ch: ch, id: who}
	active := make(chan struct{})
	go disconnectIdle(active, time.NewTimer(5*time.Minute), conn, ch, who)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		active <- struct{}{}
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func disconnectIdle(active chan struct{}, t *time.Timer, conn net.Conn, ch chan string, who string) {
	select {
	case <-active:
		t.Reset(5 * time.Minute)
	case <-t.C:
		leaving <- ch
		messages <- who + " has left"
		conn.Close()
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

// !+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
