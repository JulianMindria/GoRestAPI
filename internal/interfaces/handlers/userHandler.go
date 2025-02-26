package handlers

import (
	"GoRestAPI/internal/application/usecases"
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/interfaces/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	usecase *usecases.UserUseCase
}

func NewUserHandler(usecase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

// Createuser creates a new user
// @Summary Create a user
// @Description Add a new user to the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body entities.User true "user data"
// @Success 201 {object} entities.User
// @Failure 400 {object} string
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input struct {
		Name    string `json:"name"`
		Points  int    `json:"points"`
		Balance int    `json:"balance"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "Invalid input", err))
		return
	}

	user, err := h.usecase.CreateUser(input.Name, input.Points, input.Balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse(http.StatusInternalServerError, "Failed to create user", err))
		return
	}

	c.JSON(http.StatusCreated, common.SuccessResponse("user created successfully", user))
}

// GetAllUsers returns all users
// @Summary Get all users
// @Description Retrieve all users from the database
// @Tags users
// @Produce json
// @Success 200 {array} entities.User
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	vouchers, err := h.usecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse(http.StatusInternalServerError, "Failed to retrieve vouchers", err))
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("Vouchers retrieved successfully", vouchers))
}

// UpdateUserPoints updates a user's points
// @Summary Update User Points
// @Description Update the points of a user based on their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param points query int true "Points to be added"
// @Success 200 {object} entities.User
// @Failure 400 {object} entities.User
// @Failure 404 {object} entities.User
// @Router /users/{id}/update [put]
func (h *UserHandler) UpdateUserPoints(c *gin.Context) {
	userIDStr := c.Param("id")
	pointsStr := c.DefaultQuery("points", "0")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "Invalid user ID", err))
		return
	}

	points, err := strconv.Atoi(pointsStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "Invalid points value", err))
		return
	}

	err = h.usecase.UpdatePoints(userID, points)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse(http.StatusInternalServerError, "Failed to update points", err))
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("Points updated successfully", nil))
}

// UpdateUser updates a user's details
// @Summary Update a user
// @Description Update user details such as name, points, and balance
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body entities.User true "User data to be updated"
// @Success 200 {object} map[string]interface{} "User updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Failed to update user"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "Invalid user ID", err))
		return
	}

	var input struct {
		Name    string `json:"name"`
		Points  int    `json:"points"`
		Balance int    `json:"balance"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "Invalid input", err))
		return
	}

	user := entities.User{
		ID:      id,
		Name:    input.Name,
		Points:  input.Points,
		Balance: input.Balance,
	}

	err = h.usecase.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse(http.StatusInternalServerError, "Failed to update user", err))
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("User updated successfully", user))
}

// GetAvailableVouchersForUser retrieves vouchers a user can redeem
// @Summary Get available vouchers for a user
// @Description Fetch all vouchers that a user has enough points to redeem and that are not expired
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} entities.Voucher "List of available vouchers"
// @Failure 400 {object} map[string]interface{} "Invalid user ID"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Failed to retrieve vouchers"
// @Router /users/{id}/vouchers [get]
func (h *UserHandler) GetAvailableVouchersForUser(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "Invalid user ID", err))
		return
	}

	vouchers, err := h.usecase.GetAvailableVouchersForUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse(http.StatusInternalServerError, "Failed to retrieve vouchers", err))
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("Available vouchers retrieved successfully", vouchers))
}
