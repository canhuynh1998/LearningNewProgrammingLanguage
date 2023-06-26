package models

import (
	// "database/sql/driver"
	// "encoding/json"
	// "errors"
	"gorm.io/gorm"

	// "go-practice/backendApi/models"
)

// gorm.Model definition
// type Model struct {
//     ID        uint           `gorm:"primaryKey"`
//     CreatedAt time.Time
//     UpdatedAt time.Time
//     DeletedAt gorm.DeletedAt `gorm:"index"`
//   }

type Book struct {
    gorm.Model
    Title string 
    BookAttrs BookAttrs `gorm:"embedded"`
}




// Value make the BookAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
// func (b BookAttrs) Value() (driver.Value, error) {
//     return json.Marshal(b)
// }

// // Scan make the BookAttrs struct implement the sql.Scanner interface.
// // This method simply decodes a JSON-encoded value into the struct fields.
// func (b *BookAttrs) Scan(value interface{}) error {
//     j, ok := value.([]byte)
//     if !ok {
//         return errors.New("type assertion to []byte failed")
//     }

//     return json.Unmarshal(j, &b)
// }