package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":9000", "Server address")
	flag.Parse()

	log.Printf("starting server on %s", *addr)

	pageHtml, err := os.ReadFile("./page.html")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	http.ListenAndServe(*addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(pageHtml)
	}))
}
