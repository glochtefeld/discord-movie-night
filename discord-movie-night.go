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

	// Returns the list of servers (dev only)
	r.GET("/server", func(c *gin.Context) {

	})
	// Returns information regarding the server with the specified guild_id
	r.GET("/server/:guild_id", func(c *gin.Context) {

	})
	// Returns a list of members in a guild who have logged in
	r.GET("/server/:guild_id/members", func(c *gin.Context) {

	})
	// Returns a list of movies that the guild has suggested and watched
	r.GET("/server/:guild_id/movies", func(c *gin.Context) {

	})
	// Returns a list of scheduled events for a guild
	r.GET("/server/:guild_id/events", func(c *gin.Context) {

	})
	// Returns a list of regularly scheduled events for a guild
	r.GET("/server/:guild_id/events/recurring", func(c *gin.Context) {

	})
	// Adds a new scheduled event for a server
	r.POST("/server/:guild_id/events/new", func(c *gin.Context) {

	})
	// Returns the event details of a specified event
	r.GET("/server/:guild_id/events/:event_id", func(c *gin.Context) {

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
