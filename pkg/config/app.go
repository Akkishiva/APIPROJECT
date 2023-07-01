package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// returns a variable db which helps to interact with the DB
var (
	db *gorm.DB
)

// open a connection with the database
func Connect(){  
	d,err:=gorm.Open("mysql","root:1234@/orderdetails?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil{
		fmt.Println(d)
		panic(err)
		
	}
	db=d

}
// Returns db variable that helps us to talk to the database
func GetDB() *gorm.DB{
	return db
}