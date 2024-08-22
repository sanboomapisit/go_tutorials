package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

var db *sql.DB

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	fmt.Print(psqlInfo)
	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	db = sdb

	defer fmt.Println("close connection database !!")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Get("/product", productHandler)
	app.Get("/product/:id", getProductHandler)
	app.Post("/product", createProductHandler)
	app.Put("/product/:id", updateProductHandler)
	app.Delete("/product/:id", deleteProductHandler)
	app.Listen(":8080")
	// fmt.Println("connection database successful!!")

	// err = createProduct(&Product{Name: "go product", Price: 25})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("create successful!!")

	// product, _ := getProduct(2)
	// fmt.Println("get successful!! : ", product)

	// product, err = updateProduct(3, &Product{Name: "test go update2", Price: 232})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("update successful!!", product)

	err = deleteProduct(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete successful!!", 5)

	products, err := getProducts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("get products successful!! :", products)

}

func getProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	product, err := getProduct(productId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}

func createProductHandler(c *fiber.Ctx) error {
	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err := createProduct(p)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(p)
}

func updateProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product, err := updateProduct(productId, p)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}

func deleteProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = deleteProduct(productId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func productHandler(c *fiber.Ctx) error {
	products, err := getProducts()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(products)
}
