package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	store = map[string]string{}
	mu    sync.RWMutex
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", handlePost)
	mux.HandleFunc("GET /{uuid}", handleGet)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id := newUUID()

	mu.Lock()
	store[id] = string(body)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, id)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("uuid")

	mu.RLock()
	url, ok := store[id]
	mu.RUnlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func newUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}
