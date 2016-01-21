package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/vodolaz095/goherokuredis"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, "Redis has this error: %s!", err)
			}
		}()

		redisClient, err := goherokuredis.Init()
		if err != nil {
			panic(err)
		}
		defer func() {
			redisClient.Close()
		}()
		err = redisClient.Ping().Err()
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Redis on %s is online!", redisClient.String())
	})
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
	if err != nil {
		panic(err)
	}
}
