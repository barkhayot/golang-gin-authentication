package user_db

import (
	"log"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"	
)

var (
	Client *sql.DB
)

const (
	host     = "localhost"
  	port     = 5432
	username = "postgres"
	password = "010203"
	dbname 	 = "auth-jwt"	
)

func init() {
	dataInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port,  username, password, dbname)
	var err error
	Client, err := sql.Open("postgres", dataInfo)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	} 
	
	log.Println("Database successfully Connected")
}