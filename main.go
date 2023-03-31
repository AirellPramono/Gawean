package main

import (
	"Gawean/database"
	. "Gawean/helper"
	. "Gawean/routers"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB   *sql.DB
	err  error
	PORT = ":8080"
)

func main() {
	err = godotenv.Load("config/.env")
	CheckEnvErr(err)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname =%s sslmode=require", os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"))
	//psqlInfo := fmt.Sprintf(`host=#{os.Getenv("DB_HOST")} port=#{os.Getenv("DB_PORT")} user=#{os.Getenv("DB_USER")} password=#{os.Getenv("DB_PASSWORD")} dbname =#{os.Getenv("DB_NAME")} sslmode=disable`)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	CheckDBErr(err)
	database.DbMigrate(DB)
	defer DB.Close()
	StartServer().Run(PORT)

}
