package models

import (

	"gorm.io/gorm"
)

type BookAttrs struct {
    gorm.Model
    Picture     string 
    Description string 
    Rating      int    
}