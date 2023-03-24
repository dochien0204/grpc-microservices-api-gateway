package product

import (
	"Microservices/pkg/config"
	"Microservices/pkg/product/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	r.Group("/product")
	r.POST("/", svc.CreateProduct)
	r.GET("/find/:id", svc.FindOne)

	return svc
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}
