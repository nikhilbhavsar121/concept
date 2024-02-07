package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

const requestIdKey = 42

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIdKey).(int64)
	if !ok {
		log.Println("could not found request Id in context")

		return
	}
	log.Printf("[%d] %s", id, msg)
}

func Decorator(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIdKey, id)
		f(w, r.WithContext(ctx))
	}
}
