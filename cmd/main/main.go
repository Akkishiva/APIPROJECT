package main

import (
	"fmt"
	"github/akkishivaprakash.com/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r:=mux.NewRouter()
	routes.OrderDetails(r)
	http.Handle("/",r)
	fmt.Println("Sever started")
	log.Fatal(http.ListenAndServe("localhost:3036",r))
}