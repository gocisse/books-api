package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/mmcisse/books-api/book"
	"github.com/mmcisse/books-api/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBooks)
	app.Delete("/api/v1/book/:id", book.DeleteBooks)
	app.Put("/api/v1/book/:id", book.UpdateBooks)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "book.db")
	must(err)
	fmt.Println("Connected to database successfully ")

	// migrate the DB using the struck we declare in book package
	//this will take the Book struct and create a table in the db

	database.DBConn.AutoMigrate(&book.Book{})

	fmt.Println("Database migrated successfully ")
}

func main() {

	app := fiber.New()

	initDatabase()

	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}



func must(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
