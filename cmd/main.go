package main

import (
	"database/sql"
	"log"

	"github.com/MostafaAnas/go-ecommerce-api.git/cmd/api"
	"github.com/MostafaAnas/go-ecommerce-api.git/config"
	"github.com/MostafaAnas/go-ecommerce-api.git/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddr,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil{
		log.Fatal(err)
	}
	log.Println("DB Successfully connected")
}