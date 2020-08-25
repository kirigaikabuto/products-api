package main

import (
	"encoding/json"
	"fmt"
	"github.com/djumanoff/amqp"
)

var cfgAmqp = amqp.Config{
	Host: "localhost",
	VirtualHost: "",
	User: "",
	Password: "",
	Port: 5672,
	LogLevel: 5,
}
type Product struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Price    int64  `json:"price,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}
func main() {
	fmt.Println("Start")

	sess := amqp.NewSession(cfgAmqp)

	if err := sess.Connect(); err != nil {
		fmt.Println(err)
		return
	}
	defer sess.Close()

	var cltCfg = amqp.ClientConfig{
		//ResponseX: "response",
		//RequestX: "request",
	}

	clt, err := sess.Client(cltCfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	product := &struct {
		Id       int64  `json:"id,omitempty"`
	}{
		1,
	}
	body, err := json.Marshal(product)
	reply, err := clt.Call("products.get", amqp.Message{
		Body: body,
	})
	if err != nil{
		fmt.Println(err)
		return
	}
	newProduct := &Product{}
	err = json.Unmarshal(reply.Body, newProduct)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newProduct)
	fmt.Println("End")
}