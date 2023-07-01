package routes

import (
	"github/akkishivaprakash.com/pkg/controllers"

	"github.com/gorilla/mux"
)

var OrderDetails = func(r *mux.Router){
	r.HandleFunc("/orders",controllers.GetOrdersHandlers).Methods("GET")
	r.HandleFunc("/orders", controllers.CreateOrderHandler).Methods("POST")
	r.HandleFunc("/orders/{orderId}", controllers.GetOrderHandler).Methods("GET")
	r.HandleFunc("/orders/{orderId}", controllers.UpdateOrderHandler).Methods("PUT")
	r.HandleFunc("/orders/{orderId}", controllers.DeleteOrderHandler).Methods("DELETE")
}