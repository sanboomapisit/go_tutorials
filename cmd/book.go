package main

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func checkMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}
	// name := claims["name"].(string)
	return c.Next()
}

func getBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}

	}
	return c.SendStatus(fiber.StatusNotFound)
}
func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books[i])
		}

	}

	return c.SendStatus(fiber.StatusNotFound)
}

func deleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books = append(books[:i], books[i+1:]...) //(ตำแหน่งเริ่มต้น,ตำแหน่งที่ต้องการเอาออก)  // [:i] คือ 0 - ก่อน i ไม่เอา i
			//... แยก slice element ออก เหมือน ... ใน ่javascript เพราะ append ไม่สนับสนุนการต่อ slice กับ slide คือ append(slide,slide) ทำไม่ได้
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func testHTML(c *fiber.Ctx) error {
	// Render index template
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func getENV(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("SECRET"),
	})
}
