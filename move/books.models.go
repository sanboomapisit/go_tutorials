package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

func createBook(db *gorm.DB, book *Book) error {
	result := db.Create(book)

	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Successful!!")
	return nil
}

func getBook(db *gorm.DB, id int) *Book {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		log.Fatal("Error get book : %v", result.Error)
	}
	return &book
}

func getBooks(db *gorm.DB) []Book {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		log.Fatal("Error get book : %v", result.Error)
	}
	return books
}

func updateBook(db *gorm.DB, book *Book) error {
	result := db.Model(book).Updates(book)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Update Successful!!")
	return nil
}

func deleteBook(db *gorm.DB, id int) error {
	var book Book

	result := db.Delete(&book, id) //soft delete

	// result := db.Unscoped().Delete(&book, id) //delete row
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Delete Successful!!")
	return nil
}

func searchBook(db *gorm.DB, bookName string) []Book {
	var books []Book
	result := db.Where("name = ?", bookName).Order("price desc").Find(&books)
	if result.Error != nil {
		log.Fatal("Error search book : %v", result.Error)
	}
	return books
}
