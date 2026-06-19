package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/irenato/coverletter-hub/api/internal/config"
)

func main() {
	cfg := config.Load()
	addr := fmt.Sprintf(":%s", cfg.APIPort)
	log.Printf("starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
