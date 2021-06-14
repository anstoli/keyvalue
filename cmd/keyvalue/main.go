package main

import (
	"flag"
	"log"

	"github.com/tohast/keyvalue/http"
	"github.com/tohast/keyvalue/inmemory"
)

var port = flag.String("port", "8080", "HTTP server port")

func main()  {
	flag.Parse()
	store := inmemory.NewStore()
	err := http.Serve(*port, store)
	log.Fatal("HTTP server exited", err)
}
