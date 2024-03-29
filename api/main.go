package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	config.Load()
	r := router.CreateRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:3000", "http://localhost:3001"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := c.Handler(r)

	fmt.Printf("api online on: http://localhost:%d\n", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", config.Port), handler); err != nil {
		log.Fatal("server initialize error:", err)
	}
}
