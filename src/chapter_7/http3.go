package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	price, ok := db[item]
	if !ok {
		http.Error(w, fmt.Sprintf("no such page: %s\n", req.URL), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}

	mux := http.NewServeMux()

	// How http.HandlerFunc looks like:
	//
	// package http
	// type HandlerFunc func(w ResponseWriter, r *Request)
	//
	// func (f HandlerFunc) ServeHTTP(w ResponseWritter, r *Request) {
	// 	f(w, r)
	// }

	// This is a type conversion, not a function call.
	mux.Handle("/list", http.HandlerFunc(db.list))

	// could be simplified to:
	// mux.HandleFunc("/price", db.price)
	mux.Handle("/price", http.HandlerFunc(db.price))

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
