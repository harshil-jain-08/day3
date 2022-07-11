package Models

import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"

type Record struct {
	gorm.Model
	Roll     uint   `json:"roll"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	DOB      string `json:"dob"`_
	Address  string `json:"address"`
	Subject  string `json:"subject"`
	Marks    uint   `json:"marks"`
}

/*type Scores struct {
	ID    string `json:"id"`
	score int    `json:"marks"`
}

func (b *Scores) TableName() string {
	return "scores"
}*/

func (b *Record) TableName() string {
	return "records"
}