package models

import (
	"github/akkishivaprakash.com/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Order represents the structure of an order
type Address struct {
    Street     string `json:"street"`
    City       string `json:"city"`
    State      string `json:"state"`
    PostalCode string `json:"postal_code"`
}

type Customer struct {

    Name    string  `json:"name"`
    Email   string  `json:"email"`
    Address1 Address  `gorm:"embedded" json:"Address"`
}

type Order struct {
    gorm.Model
    // OrderID     int      `gorm:"uniqueIndex" json:"order_id"`
    Customer1    Customer `gorm:"embedded" json:"customer"`
    ProductName string   `json:"product_name"`
    Quantity    int      `json:"quantity"`
    OrderDate   string   `json:"order_date"`
    Priority    string   `json:"priority"`
}
// initalize the database
func init(){
	config.Connect()
	db=config.GetDB()
	db.AutoMigrate(&Order{})
}

func (o *Order) CreateOrder() *Order{
	db.NewRecord(o)
	db.Create(&o)
	return o
}

func GetAllOrders() []Order{
	var Orders []Order
	db.Find(&Orders)
	return Orders
}

func GetOrderById(Id int64)(*Order,*gorm.DB){
	var getOrder Order
	db:=db.Where("ID=?",Id).Find(&getOrder)
	return &getOrder,db

}

func DeleteBook(ID int64)Order{
	var order Order
	db.Where("ID=?",ID).Delete(order)
	return order 
}