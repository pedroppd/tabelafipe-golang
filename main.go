package main

import (
	"fmt"
	"log"
	"net/http"
	"tabela-fipe-golang/router"
)

func main() {
	fmt.Println("API is running...")
	router := router.Generate()
	log.Fatal(http.ListenAndServe(":5000", router))
}
