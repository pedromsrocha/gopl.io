// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//!+main

func main() {
	db := &database{m: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	m  map[string]dollars
	mu sync.Mutex
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	for item, price := range db.m {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
	db.mu.Unlock()
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db.mu.Lock()
	defer db.mu.Unlock()
	if price, ok := db.m[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	priceV, err := strconv.ParseFloat(price, 32)
	if err != nil {
		http.Error(w, "Price is not a valid number", http.StatusBadRequest)
		return
	}
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, ok := db.m[item]; ok {
		http.Error(w, fmt.Sprintf("Item %s already exists", item), http.StatusBadRequest)
		return
	}
	db.m[item] = dollars(priceV)
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, ok := db.m[item]; !ok {
		http.Error(w, fmt.Sprintf("Item %s does not exist", item), http.StatusBadRequest)
		return
	}
	delete(db.m, item)
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	priceV, err := strconv.ParseFloat(price, 32)
	if err != nil {
		http.Error(w, "Price is not a valid number", http.StatusBadRequest)
		return
	}
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, ok := db.m[item]; !ok {
		http.Error(w, fmt.Sprintf("Item %s does not exist", item), http.StatusBadRequest)
		return
	}
	db.m[item] = dollars(priceV)
}
