package main

import (
	"fmt"
	"learncontext/log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", log.Decorator(handler))
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println(ctx, "handle started")
	defer log.Println(ctx, "handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "hello buddy")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInsufficientStorage)
	}

}
