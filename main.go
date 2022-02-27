package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // initialise this library but don't use it
	"log"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
	table    = "users"
	sslmode  = "disable"
	port     = 5432
)

func main() {
	fmt.Println("hello")
	connString := fmt.Sprintf("user=%s password='%s' dbname=%s port=%d sslmode=%s", user, password, dbname, port, sslmode)
	log.Println(connString)

	// this by itself does not open conn to db !
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalln(err)
	}
	// close the db connection at the end
	defer db.Close()

	// verify connection to our db by forcig our code to make a connection to db !!
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("db connected !!")

	var insertId int
	// insert using hardcode
	insertStatement := `INSERT INTO users (age, email, first_name, last_name)
	VALUES (23, 'oneal@shaquille.nba', 'shaq', 'oneal') RETURNING id`
	err = db.QueryRow(insertStatement).Scan(&insertId)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("last insert id : %d ", insertId)

	// insert using string interpolation
	insertStatmentInterpolate := `INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4) RETURNING id`
	// by doing this, `database/sql` pkg also sanitises our sql string before executing our commands !
	// so it is better to just do this instead of using fmt.sprintf to create our query string !
	err = db.QueryRow(insertStatmentInterpolate, 40, "steve.nash@nba.com", "steve", "nash").Scan(&insertId)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("last insert id : %d ", insertId)

}
