package main

import (
	"Microservices/pkg/auth"
	"Microservices/pkg/config"
	"Microservices/pkg/order"
	"Microservices/pkg/product"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config:", err)
	}

	r := gin.Default()

	auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c)
	order.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
