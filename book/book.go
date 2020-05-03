package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/mmcisse/books-api/database"
)

type Book struct {
	gorm.Model
	Title  string  `json:"title`
	Author string  `json:"author`
	Isbn   float64 `json:isbn`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)

}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}
func NewBooks(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)
	err := c.BodyParser(book)
	if err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}
func DeleteBooks(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found with the given ID")
		return
	}
	db.Delete(book)
	c.JSON(book)

}
func UpdateBooks(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found with the given ID")
		return
	}
	db.Create(&book)
	err := c.BodyParser(book)
	if err != nil {
		c.Status(503).Send(err)
		return
	}

	c.JSON(book)
}
