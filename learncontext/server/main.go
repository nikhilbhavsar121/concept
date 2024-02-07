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

//using context.withTimeout(ctx,time) give cancel func when you call that ctx.Done() trigger and perform your operation.
//if client stop nor not able to receive the respond the using context we can stop there so n need to do any operation after that.
// We can set value using context.withValue(ctx,key,value) and get value using ctx.Value
