package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	var Cards = [...]string{"大吉", "吉", "中吉", "小吉", "凶"}

	http.HandleFunc("/", (func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", Cards[rand.Intn(5)])
	}))

	http.ListenAndServe(":80", nil)
}
