package main

import (
	handler "RajendranDinesh/goProxy/api"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/proxy", handler.Handler)

	handler := cors.AllowAll().Handler(mux)

	port := ":3000"
	fmt.Println("Proxy server running on port" + port)
	log.Fatal(http.ListenAndServe(port, handler))
}
