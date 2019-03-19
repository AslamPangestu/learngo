package db;

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

func Connect()  {
	opts:=&pg.Options{
		User:		"admin",
		Password:	"12345",
		Addr:		"localhost:5432",
	}
	var db *pg.DB=pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to DB.\n")
		os.Exit(100)
	}
	log.Printf("Succes to connect to DB.\n")
	closeErr := db.Close()
	if closeErr != nil{
		log.Printf("Error while close connection db. Reason : %v\n",closeErr)
		os.Exit(100)
	}
	log.Printf("Connection close.\n");
	return
}