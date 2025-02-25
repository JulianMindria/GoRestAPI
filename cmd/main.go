package main

import (
	"GoRestAPI/internal/application/usecases"
	"GoRestAPI/internal/infrastructure"
	"GoRestAPI/internal/infrastructure/repositories"
	"GoRestAPI/internal/interfaces/routes" // Import the new routes package
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "GoRestAPI/docs"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GoRestAPI Swagger Example API
// @version 1.0
// @description This is a simple Go REST API with Swagger documentation.
// @host localhost:8080
// @BasePath /api
func main() {
	db, err := infrastructure.ConnectDatabase()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	voucherRepo := repositories.NewVoucherRepository(db)
	voucherUseCase := usecases.NewVoucherUseCase(voucherRepo)

	userRepo := repositories.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)

	brandRepo := repositories.NewBrandRepository(db)
	brandUseCase := usecases.NewBrandUseCase(brandRepo)

	productRepo := repositories.NewProductRepository(db)
	productUseCase := usecases.NewProductUseCase(productRepo)

	transactionRepo := repositories.NewTransactionRepository(db)
	transactionUseCase := usecases.NewTransactionUseCase(transactionRepo)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	routes.SetUpRoutes(r, voucherUseCase, userUseCase, transactionUseCase, productUseCase, brandUseCase)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
