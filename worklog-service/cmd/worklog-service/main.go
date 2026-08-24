package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nhom2/worklog-service/internal/httpapi"
	"github.com/nhom2/worklog-service/internal/store"
)

func main() {
	addr := os.Getenv("WORKLOG_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	// The current experiment baseline intentionally uses an in-memory store.
	handler := httpapi.NewHandler(store.NewSeededWorklogStore())
	log.Printf("worklog-service listening on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
