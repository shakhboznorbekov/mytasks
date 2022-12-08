package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "2711"
	dbname   = "korzinka"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	update := `
	update user1 set first_name =$2,last_name=$3,nickname=$4
	where id = $1`

	_, err = db.Exec(update, "9b614eb0-a2bc-41b1-92cb-c87765446b89", "NewFirst", "NewLast", "hello")
	if err != nil {
		panic(err)
	}

	delete := `delete from user1 where id=$1`
	_, err = db.Exec(delete, "1254b3d3-cd20-4b1e-ac71-ccd2c6272e30")
	if err != nil {
		panic(err)
	}

}
