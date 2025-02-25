package routes

import (
	"GoRestAPI/internal/application/usecases"
	"GoRestAPI/internal/interfaces/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine,
	voucherUseCase *usecases.VoucherUseCase,
	userUseCase *usecases.UserUseCase,
	transactionUseCase *usecases.TransactionUseCase,
	productUseCase *usecases.ProductUseCase,
	brandUseCase *usecases.BrandUseCase) {

	// Initialize handlers
	voucherHandler := handlers.NewVoucherHandler(voucherUseCase)
	userHandler := handlers.NewUserHandler(userUseCase)
	transactionHandler := handlers.NewTransactionHandler(transactionUseCase)
	productHandler := handlers.NewProductHandler(productUseCase)
	brandHandler := handlers.NewBrandHandler(brandUseCase)

	api := r.Group("/api")
	{
		api.POST("/vouchers", voucherHandler.CreateVoucher)
		api.GET("/vouchers", voucherHandler.GetAllVouchers)

		api.POST("/users", userHandler.CreateUser)
		api.PUT("/users/:id/update", userHandler.UpdateUserPoints)
		api.GET("/users", userHandler.GetAllUsers)

		api.POST("/transactions", transactionHandler.CreateTransaction)
		api.GET("/transactions/:id", transactionHandler.GetTransactionByID)
		api.GET("/transactions", transactionHandler.GetAllTransactions)
		api.DELETE("/transactions/:id", transactionHandler.DeleteTransaction)

		api.POST("/products", productHandler.CreateProduct)
		api.GET("/products", productHandler.GetAllProducts)
		api.GET("/products/:id", productHandler.GetProductByID)
		api.DELETE("/products/:id", productHandler.DeleteProduct)

		api.POST("/brands", brandHandler.CreateBrand)
		api.GET("/brands", brandHandler.GetAllBrands)
		api.GET("/brands/:id", brandHandler.GetBrandByID)
		api.DELETE("/brands/:id", brandHandler.DeleteBrand)
	}
}
