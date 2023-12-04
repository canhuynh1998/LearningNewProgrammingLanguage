package services

import (
	"go-practice/backendApi/databases"
	"go-practice/backendApi/models"
)

var BookRepository = databases.GetDBInstance().DB.Model(&models.Book{})

func GetBookFromId(id string) (models.Book) {
	targetBook := models.Book{}
	var book models.Book  
	BookRepository.Find(&targetBook, id).Scan(&book)


	return book
}

