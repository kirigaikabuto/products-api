package main

import (
	"fmt"
	"github.com/djumanoff/amqp"
	"github.com/kirigaikabuto/products"
	"log"
)

var cfg = amqp.Config{
	Host: "localhost",
	VirtualHost: "",
	User: "",
	Password: "",
	Port: 5672,
	LogLevel: 5,
}

var srvCfg = amqp.ServerConfig{
	//ResponseX: "response",
	//RequestX: "request",
}

func main() {
	fmt.Println("Start")

	sess := amqp.NewSession(cfg)

	if err := sess.Connect(); err != nil {
		fmt.Println(err)
		return
	}
	defer sess.Close()

	srv, err := sess.Server(srvCfg)
	if err != nil {
		fmt.Println(err)
		return
	}
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
	productsAmqpEndpoints := products.NewAmqpEndpointFactory(productService)
	srv.Endpoint("products.get",productsAmqpEndpoints.GetProductByIdAMQPEndpoint())
	if err := srv.Start(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("End")
}