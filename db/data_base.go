package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect_db() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/ydays")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := "INSERT INTO `user` (`email`, `nom`, `prenom`, `password`)VALUES (?,?,?,?)"
	insertResult, err := db.ExecContext(context.Background(), query, "aurelie.lao@ynov.com", "Lao", "Aurelie", "aurelie_123")
	if err != nil {
		log.Fatalf("impossible insert teacher: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}
	log.Printf("inserted id: %d", id)
	fmt.Println("Success!")
}

func Create_db(email string, nom string, prenom string, password string) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/ydays")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := "INSERT INTO `user` (`email`, `nom`, `prenom`, `password`)VALUES (?,?,?,?)"
	insertResult, err := db.ExecContext(context.Background(), query, email, nom, prenom, password)
	if err != nil {
		log.Fatalf("impossible insert teacher: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}
	log.Printf("inserted id: %d", id)
}
