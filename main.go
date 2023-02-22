package main

import (
	"fmt"
	"log"
	"main/router"
	"net/http"
)

func main() {
	fmt.Println("Mongo db apis")

	r := router.Router()
	fmt.Println("starting server...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("server listening at 4000")
}
