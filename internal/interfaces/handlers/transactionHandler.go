package handlers

import (
	"GoRestAPI/internal/application/dto"
	"GoRestAPI/internal/application/usecases"
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/interfaces/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	transactionUseCase *usecases.TransactionUseCase
}

func NewTransactionHandler(uc *usecases.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{transactionUseCase: uc}
}

// @Summary Create a new transaction
// @Description Create a new transaction and store it in the database
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body entities.Transaction true "Transaction Data"
// @Success 200 {object} entities.Transaction
// @Failure 400 {object} entities.Transaction
// @Router /transactions [post]
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transaction entities.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.transactionUseCase.CreateTransaction(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	transactionDTO := dto.NewTransactionDTO(&transaction)

	c.JSON(http.StatusCreated, common.SuccessResponse("transaction created successfully", transactionDTO))
}

// @Summary Get a transaction by ID
// @Description Get transaction details by its ID
// @Tags transactions
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} entities.Transaction
// @Failure 404 {object} entities.Transaction
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetTransactionByID(c *gin.Context) {
	transactionIDStr := c.Param("id")
	transactionID, err := uuid.Parse(transactionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := h.transactionUseCase.GetTransactionByID(transactionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	transactionDTO := dto.NewTransactionDTO(transaction)

	c.JSON(http.StatusOK, common.SuccessResponse("transaction retrieved successfully", transactionDTO))
}

// @Summary Get all transactions
// @Description Get a list of all transactions
// @Tags transactions
// @Produce json
// @Success 200 {array} entities.Transaction
// @Failure 500 {object} entities.Transaction
// @Router /transactions [get]
func (h *TransactionHandler) GetAllTransactions(c *gin.Context) {
	transactions, err := h.transactionUseCase.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("transaction retrieved successfully", transactions))
}

// @Summary Delete a transaction
// @Description Delete a transaction by its ID
// @Tags transactions
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} entities.Transaction
// @Failure 404 {object} entities.Transaction
// @Router /transactions/{id} [delete]
func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	transactionIDStr := c.Param("id")
	transactionID, err := uuid.Parse(transactionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	err = h.transactionUseCase.DeleteTransaction(transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("transaction deleted successfully", transactionID))
}

// @Summary Get Transaction Details by User
// @Description Retrieve all transactions for a specific user
// @Tags transactions
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {array} entities.Transaction
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /transactions/user/{user_id} [get]
func (h *TransactionHandler) GetTransactionDetailByUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	transactions, err := h.transactionUseCase.GetTransactionDetailByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}

	var transactionDTOs []dto.TransactionDetailDTO
	for _, transaction := range transactions {
		transactionDTOs = append(transactionDTOs, dto.NewTransactionDetailDTO(&transaction))
	}

	c.JSON(http.StatusOK, common.SuccessResponse("Transactions retrieved successfully", transactionDTOs))
}
