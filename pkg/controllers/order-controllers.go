package controllers

import (
	"encoding/json"
	"fmt"
	"github/akkishivaprakash.com/pkg/models"
	"github/akkishivaprakash.com/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewOrder models.Order

func GetOrdersHandlers(w http.ResponseWriter,r *http.Request){
	// newOrders:=models.GetAllOrders()
	// res,_:=json.Marshal(newOrders)
	// w.Header().Set("Content-Type","pkglication/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
	w.Header().Set("content-Type","application/json")
	json.NewEncoder(w).Encode(models.GetAllOrders())

}

func CreateOrderHandler(w http.ResponseWriter,r *http.Request){
	CreateOrder:=&models.Order{}
	utils.ParseBody(r,CreateOrder)
	o:=CreateOrder.CreateOrder()
	res,_:=json.Marshal(o)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
// get the single order
func GetOrderHandler(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	orderId:=vars["orderId"]
	ID,err:=strconv.ParseInt(orderId,10,64)
	if err !=nil{
		fmt.Println(orderId)
		fmt.Println("error while parsing")
		fmt.Println(err)
		fmt.Println(vars)
	}
	orderDetails,_:=models.GetOrderById(ID)
	res,_:=json.Marshal(orderDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}

func UpdateOrderHandler(w http.ResponseWriter,r *http.Request){
        updateOrder:=&models.Order{}
		utils.ParseBody(r,updateOrder)
		vars:=mux.Vars(r)
		orderId:=vars["orderId"]
		ID ,err:=strconv.ParseInt(orderId,0,0)
		if err!=nil{
			fmt.Println("error while parsing")
			fmt.Println(err)
		} 
		orderDetails,db:=models.GetOrderById(ID)
		// if updateOrder.ID!=0{
		// 	orderDetails.ID=updateOrder.ID
		// }
		if updateOrder.Customer1.Name!=""{
			orderDetails.Customer1.Name=updateOrder.Customer1.Name

		}
		if updateOrder.Customer1.Email!=""{
			orderDetails.Customer1.Email=updateOrder.Customer1.Email

		}
		if updateOrder.Customer1.Address1.Street!=""{
			orderDetails.Customer1.Address1.Street=updateOrder.Customer1.Address1.Street

		}
		if updateOrder.Customer1.Address1.City!=""{
			orderDetails.Customer1.Address1.City=updateOrder.Customer1.Address1.City

		}
		if updateOrder.Customer1.Address1.State!=""{
			orderDetails.Customer1.Address1.State=updateOrder.Customer1.Address1.State

		}
		if updateOrder.Customer1.Address1.PostalCode!=""{
			orderDetails.Customer1.Address1.PostalCode=updateOrder.Customer1.Address1.PostalCode

		}
		if updateOrder.ProductName!=""{
			orderDetails.ProductName=updateOrder.ProductName

		}
		if updateOrder.Quantity!=0{
			orderDetails.Quantity=updateOrder.Quantity

		}
		if updateOrder.OrderDate!=""{
			orderDetails.OrderDate=updateOrder.OrderDate

		}
		if updateOrder.Priority!=""{
			orderDetails.Priority=updateOrder.Priority

		}
		db.Save(&orderDetails)
		res,_:=json.Marshal(orderDetails)
		w.Header().Set("Content-Type","pkglication.json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
}


func DeleteOrderHandler(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	orderId:=vars["orderId"]
	ID,err:=strconv.ParseInt(orderId,0,0)
	if err !=nil{
		fmt.Println(vars)
		fmt.Println("error while parsing")
		fmt.Println(err)
	}
	order :=models.DeleteBook(ID)
	res,_:=json.Marshal(order)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

