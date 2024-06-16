package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

func RedisClient() {

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "redis", // no password set
		DB:       0,       // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pong)

	// save in redis
	err = rdb.Set(ctx, "abc", "123", 0).Err()
	if err != nil {
		log.Println("error to save in redis:", err)
	}

	// get from redis
	val, err := rdb.Get(ctx, "abc").Result()
	if err != nil {
		log.Println("error to get from redis:", err)
	}
	log.Println("abc", val)

	// get create time of key
	val, err = rdb.TTL(ctx, "abc").Result()
	if err != nil {
		log.Println("error to get TTL from redis:", err)
	}
}

func main() {

	RedisClient()

	// Create Middlware Function to count request from IP or Token
	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Middleware")

			log.Println("IP Requester:", strings.Split(r.RemoteAddr, ":")[0])

			next.ServeHTTP(w, r)
		})
	}

	// Create ServeMux to handle request
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!\n"))
	})

	// Chain Middleware with ServeMux
	chainRequest := middleware(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server with ServeMux
	log.Println("Start WebServer Port", port)
	if err := http.ListenAndServe(":"+port, chainRequest); err != nil {
		log.Fatal(err)
	}
}
