package main 

import (
	"net/http"
	"todo-go/internal/http"
	"log"
)

func main() {
	http.RegisterRoutes()
	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
