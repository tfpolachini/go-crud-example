package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tfpolachini/go-crud-example/application/controller"
	"github.com/tfpolachini/go-crud-example/domain/service"
	"github.com/tfpolachini/go-crud-example/infrastructure/database"
	"github.com/tfpolachini/go-crud-example/infrastructure/repository"
)

func main() {

	start := time.Now()

	conn := database.Connect()

	defer conn.Close()

	router := gin.Default()

	router.POST("/products", MakeCreateProductHandler(conn))

	fmt.Printf("Application started in ~%3.3f seconds\n", time.Since(start).Seconds())

	router.Run(":8080")
}

func MakeCreateProductHandler(conn *sql.DB) func(*gin.Context) {
	repo := repository.NewProductRepository(conn)
	svc := service.NewProductService(repo)
	return controller.NewProductController(svc).CreateProduct
}
