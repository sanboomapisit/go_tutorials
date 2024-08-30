package main

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt/v4"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// const (
// 	host     = "localhost"  // or the Docker service name if running in another container
// 	port     = 5432         // default PostgreSQL port
// 	user     = "myuser"     // as defined in docker-compose.yml
// 	password = "mypassword" // as defined in docker-compose.yml
// 	dbname   = "mydatabase" // as defined in docker-compose.yml
// )

// func authRequired(c *fiber.Ctx) error {
// 	cookie := c.Cookies("jwt")
// 	jwtSecretKey := "mySecretKey"
// 	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(jwtSecretKey), nil
// 	})

// 	if err != nil || !token.Valid {
// 		return c.SendStatus(fiber.StatusUnauthorized)
// 	}
// 	claim := token.Claims.(jwt.MapClaims)

// 	fmt.Println(claim["user_id"])
// 	return c.Next()
// }

// func main() {
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	newLogger := logger.New(
// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
// 		logger.Config{
// 			SlowThreshold: time.Second, // Slow SQL threshold
// 			LogLevel:      logger.Info, // Log level
// 			Colorful:      true,        // Enable color
// 		},
// 	)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: newLogger,
// 	})
// 	if err != nil {
// 		panic("failed to connect to database")
// 	}

// 	// db.Migrator().DropColumn(&Book{}, "name")
// 	db.AutoMigrate(&Book{}, &User{}) // ไม่ลบ column ทำได้แค่เพิ่มเท่านั้น
// 	// db.Migrator().RenameColumn(&Book{}, "name","last_name")
// 	fmt.Println("connect successful")
// 	// newBook := Book{Name: "the Cybersecurity",
// 	// 	Author:      "sun apisit",
// 	// 	Description: "1 hard basic cyber security 101",
// 	// 	Price:       40}

// 	// createBook(db, &newBook)
// 	// currentBook := getBook(db, 1)
// 	// currentBook.Name = "new book name3"
// 	// currentBook.Price = 140
// 	// updateBook(db, currentBook)
// 	// fmt.Println(*currentBook)
// 	// deleteBook(db, 1)
// 	// createBook := searchBook(db, "the Cybersecurity")
// 	// for _, book := range createBook {
// 	// 	fmt.Println(book.ID, book.Name, book.Author, book.Price)
// 	// }
// 	// fmt.Println(createBook)
// 	app := fiber.New()
// 	app.Use("/books", authRequired)
// 	app.Get("/books", func(c *fiber.Ctx) error {
// 		return c.JSON(getBooks(db))
// 	})
// 	app.Get("/books/:id", func(c *fiber.Ctx) error {
// 		id, err := strconv.Atoi(c.Params("id"))
// 		if err != nil {
// 			return c.SendStatus(fiber.StatusBadRequest)
// 		}
// 		book := getBook(db, id)
// 		return c.JSON(book)
// 	})
// 	app.Post("/books", func(c *fiber.Ctx) error {
// 		book := new(Book)
// 		if err := c.BodyParser(book); err != nil {
// 			return c.SendStatus(fiber.StatusBadRequest)
// 		}
// 		err = createBook(db, book)
// 		if err != nil {
// 			return c.SendStatus(fiber.StatusBadRequest)
// 		}
// 		return c.JSON(fiber.Map{"message": "create book successful!!"})
// 	})
// 	app.Put("/books/:id", func(c *fiber.Ctx) error {
// 		id, err := strconv.Atoi(c.Params("id"))
// 		if err != nil {
// 			return c.SendStatus(fiber.StatusBadRequest)
// 		}

// 		book := new(Book)
// 		if err := c.BodyParser(book); err != nil {
// 			return c.SendStatus(fiber.StatusBadRequest)
// 		}

// 		book.ID = uint(id)

// 		err = updateBook(db, book)
// 		if err != nil {
// 			return c.SendStatus(fiber.StatusBadRequest)
// 		}
// 		return c.JSON(fiber.Map{"message": "update book successful!!"})
// 	})
// 	app.Delete("/books/:id", func(c *fiber.Ctx) error {
// 		id, err := strconv.Atoi(c.Params("id"))
// 		if err != nil {
// 			return c.SendStatus(fiber.StatusBadRequest)
// 		}

// 		err = deleteBook(db, id)
// 		if err != nil {
// 			return c.SendStatus(fiber.StatusBadRequest)
// 		}
// 		return c.JSON(fiber.Map{"message": "delete book successful!!"})
// 	})
// 	app.Post("/register", func(c *fiber.Ctx) error {
// 		user := new(User)
// 		if err := c.BodyParser(user); err != nil {
// 			return c.SendStatus(fiber.StatusBadRequest)
// 		}
// 		err = createUser(db, user)
// 		if err != nil {
// 			return c.SendStatus(fiber.StatusBadRequest)
// 		}
// 		return c.JSON(fiber.Map{"message": "register user successful!!"})
// 	})

// 	app.Post("/login", func(c *fiber.Ctx) error {
// 		user := new(User)
// 		if err := c.BodyParser(user); err != nil {
// 			return c.SendStatus(fiber.StatusBadRequest)
// 		}
// 		token, err := loginUser(db, user)
// 		if err != nil {
// 			return c.SendStatus(fiber.StatusUnauthorized)
// 		}

// 		c.Cookie(&fiber.Cookie{
// 			Name:     "jwt",
// 			Value:    token,
// 			Expires:  time.Now().Add(time.Hour * 72),
// 			HTTPOnly: true,
// 		})
// 		return c.JSON(fiber.Map{
// 			"message": "login successful!!",
// 		})
// 	})

// 	app.Listen(":8080")
// }
