package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	r := router.CreateRouter()

	fmt.Printf("api online on: http://localhost:%d\n", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", config.Port), r); err != nil {
		log.Fatal("server initialize error:", err)
	}
}
