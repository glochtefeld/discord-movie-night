package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

type Env struct {
	db *sql.DB
}

func setupRouter(env *Env) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}

func setupDB(dbinfo string) *sql.DB {
	fmt.Print(dbinfo)
	fmt.Print("\n")
	conn, err := pq.NewConnector(dbinfo)
	if err != nil {
		fmt.Print(err)
		log.Fatal("Could not connect to database")
	}
	db := sql.OpenDB(conn)
	return db
}

func main() {
	denvErr := godotenv.Load()
	if denvErr != nil {
		log.Fatal("Error loading .env file")
	}
	dbinfo := fmt.Sprintf("host=localhost port=5432 dbname=%s user=%s password=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), "dmn")
	db := setupDB(dbinfo)
	defer db.Close()
	env := &Env{db: db}

	r := setupRouter(env)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
