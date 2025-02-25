package handlers

import (
	"GoRestAPI/internal/application/usecases"
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
		ID     uuid.UUID `json:"id"`
		Name   string    `json:"name"`
		Points int       `json:"points"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "Invalid input", err))
		return
	}

	user, err := h.usecase.CreateUser(input.Name, input.Points)
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
