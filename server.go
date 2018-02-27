package main

import (
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

// Team schema
type Team struct {
	gorm.Model
	SlackID  string
	Domain   string
	Channels []Channel
}

// Channel schema
type Channel struct {
	gorm.Model
	SlackID string
	Name    string
	TeamID  uint
}
type Response struct {
	Text string `json:"text"`
}

func createTeam(c echo.Context) error {
	r := &Response{Text: "Hello World!"}
	return c.JSON(http.StatusOK, r)
}

func main() {
	// db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=slack_meetup_development password=postgres")
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// db.AutoMigrate(&Team{}, &Channel{})

	e := echo.New()
	e.POST("/slack/slash_commands", createTeam)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
