package main

import (
	"fmt"
	"log"
	"net/http"
	"task1/routes"
)

func main(){
	fmt.Println("mongo db")
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":6000", r))
	fmt.Println("Listening on port 6000")
}