package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Env struct {
	db *gorm.DB
}

func generateState(n int) (string, error) {
	data := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func setupRouter(env *Env) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	sessionNames := []string{"a", "b"}
	router.Use(sessions.SessionsMany(sessionNames, store))
	router.LoadHTMLFiles("build/index.html")

	router.SetTrustedProxies(nil)

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/login", func(c *gin.Context) {
		conf := &oauth2.Config{
			ClientID:     os.Getenv("DISCORD_CLIENT_ID"),
			ClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
			Scopes:       []string{"bot", "identify"},
			RedirectURL:  "http://localhost:8080/auth/callback",
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://discord.com/oauth2/authorize",
				TokenURL: "https://discord.com/api/oauth2/token",
			},
		}
		/* 		state, err := generateState(128)
		   		if err != nil {
		   			log.Fatal(err)
		   		} */
		url := conf.AuthCodeURL("hello", oauth2.AccessTypeOffline)
		fmt.Printf("Visit the URL for the auth dialog: %v", url)

		tok, err := conf.Exchange(c, c.Query("code"))
		if err != nil {
			log.Fatal(err)
		}

		client := conf.Client(c, tok)
		fmt.Printf("%+v\n", client)
	})

	router.POST("/auth/callback", func(c *gin.Context) {
		fmt.Printf("Callback called back")
	})

	// Returns the list of servers (dev only)
	router.GET("/server", func(c *gin.Context) {

	})
	// Returns information regarding the server with the specified guild_id
	router.GET("/server/:guild_id", func(c *gin.Context) {

	})
	// Returns a list of members in a guild who have logged in
	router.GET("/server/:guild_id/members", func(c *gin.Context) {

	})
	// Returns a list of movies that the guild has suggested and watched
	router.GET("/server/:guild_id/movies", func(c *gin.Context) {

	})
	// Returns a list of scheduled events for a guild
	router.GET("/server/:guild_id/events", func(c *gin.Context) {

	})
	// Returns a list of regularly scheduled events for a guild
	router.GET("/server/:guild_id/events/recurring", func(c *gin.Context) {

	})
	// Adds a new scheduled event for a server
	router.POST("/server/:guild_id/events/new", func(c *gin.Context) {

	})
	// Returns the event details of a specified event
	router.GET("/server/:guild_id/events/:event_id", func(c *gin.Context) {

	})
	return router
}

func setupDB(dbinfo string) *gorm.DB {
	// fmt.Print(dbinfo)
	// fmt.Print("\n")
	db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
		log.Fatal("Could not connect to database")
	}
	return db
}

/* func testDB(env *Env) {
	guild := Guild{Discord_ID: "02837", Name: "The Bastard Brigade"}
	env.db.Create(&guild)
	user := User{Discord_ID: "0123", Nickname: "Libnits"}
	env.db.Create(&user)
	gm := GuildMember{User: user, Guild: guild}
	env.db.Create(&gm)

	var users []struct {
		Nickname   string
		ServerName string
	}
	var users1 []struct {
		Nickname   string
		ServerName string
	}

	s2 := time.Now()
	res1 := env.db.Model(&GuildMember{}).
		Select("users.nickname, guilds.name as server_name").
		Joins("inner join guilds on guild_members.guild_id=guilds.id").
		Joins("inner join users on guild_members.user_id=users.id").
		Find(&users1)
	r2Time := time.Since(s2)

	start := time.Now()
	res := env.db.Raw(`select nickname, guilds.name as server_name
		from guild_members gm
		inner join users on gm.user_id=users.id
		inner join guilds on gm.guild_id=guilds.id`).Find(&users)
	resTime := time.Since(start)

	if res.Error == nil {
		fmt.Printf("Raw took %s; %+v\n", resTime, users)
	}
	if res1.Error == nil {
		fmt.Printf("Functiony took %s; %+v\n", r2Time, users1)
	}
}
*/

func main() {
	denvErr := godotenv.Load()
	if denvErr != nil {
		log.Fatal("Error loading .env file")
	}
	dbinfo := fmt.Sprintf("host=localhost port=5432 dbname=%s user=%s password=%s", "dmn", os.Getenv("DB_USER"), os.Getenv("DB_PASS"))
	db := setupDB(dbinfo)
	db.AutoMigrate(&Guild{}, &User{}, &GuildMember{}, &Movie{}, &GuildMovie{}, &EventLoc{}, &Event{})
	env := &Env{db: db}

	router := setupRouter(env)
	router.Run(":8080")
}
