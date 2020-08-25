package main

import (
	"github.com/gorilla/mux"
	"github.com/kirigaikabuto/products"
	"log"
	"net/http"
)
var (
	PORT string = "8000"
)

func main() {
	postgreConf := products.Config{
		Host:             "localhost",
		User:             "kirito",
		Password:         "passanya",
		Port:             5432,
		Database:         "crm",
		ConnectionString: "",
		Params:           "sslmode=disable",
	}
	productStore, err := products.NewPostgreStore(postgreConf)
	if err != nil {
		log.Fatal(err)
	}
	productService := products.NewProductService(productStore)
	productHttpEndpoints := products.NewHttpEndpoints(productService)
	router:=mux.NewRouter()
	router.Methods("POST").Path("/products").HandlerFunc(productHttpEndpoints.CreateProductEndpoint())
	router.Methods("GET").Path("/products").HandlerFunc(productHttpEndpoints.ListProductsEndpoint())
	router.Methods("PUT").Path("/products/{id}").HandlerFunc(productHttpEndpoints.UpdateProductEndpoint("id"))
	router.Methods("GET").Path("/products/{id}").HandlerFunc(productHttpEndpoints.GetProductByIdEndpoint("id"))
	log.Println("Server is running on port "+PORT)
	err = http.ListenAndServe(":"+PORT,router)
	if err != nil{
		log.Fatal(err)
	}
}
