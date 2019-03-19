package main;

import (
	"log"

	db "./db"
)


func main()  {
	log.Printf("Testing\n")
	db.Connect()
}