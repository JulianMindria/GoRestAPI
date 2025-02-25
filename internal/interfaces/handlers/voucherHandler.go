package handlers

import (
	"GoRestAPI/internal/application/usecases"
	"GoRestAPI/internal/interfaces/common"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VoucherHandler struct {
	usecase *usecases.VoucherUseCase
}

func NewVoucherHandler(usecase *usecases.VoucherUseCase) *VoucherHandler {
	return &VoucherHandler{usecase: usecase}
}

// CreateVoucher creates a new voucher
// @Summary Create a voucher
// @Description Add a new voucher to the database
// @Tags vouchers
// @Accept json
// @Produce json
// @Param voucher body entities.Voucher true "Voucher data"
// @Success 201 {object} entities.Voucher
// @Failure 400 {object} string
// @Router /vouchers [post]
func (h *VoucherHandler) CreateVoucher(c *gin.Context) {
	var input struct {
		Code        string    `json:"code"`
		CostInPoint int       `json:"cost_in_point"`
		Expiration  time.Time `json:"expiration"`
		Type        string    `json:"type"`
		Value       float64   `json:"value"`
		BrandID     uuid.UUID `json:"brand_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "Invalid input", err))
		return
	}

	voucher, err := h.usecase.CreateVoucher(input.Code, input.CostInPoint, input.Expiration, input.Type, input.Value, input.BrandID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse(http.StatusInternalServerError, "Failed to create voucher", err))
		return
	}

	c.JSON(http.StatusCreated, common.SuccessResponse("Voucher created successfully", voucher))
}

// GetAllVouchers returns all vouchers
// @Summary Get all vouchers
// @Description Retrieve all vouchers from the database
// @Tags vouchers
// @Produce json
// @Success 200 {array} entities.Voucher
// @Router /vouchers [get]
func (h *VoucherHandler) GetAllVouchers(c *gin.Context) {
	vouchers, err := h.usecase.GetAllVouchers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse(http.StatusInternalServerError, "Failed to retrieve vouchers", err))
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("Vouchers retrieved successfully", vouchers))
}
