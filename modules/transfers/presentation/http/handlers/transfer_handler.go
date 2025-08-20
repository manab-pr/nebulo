package handlers

import (
	"net/http"

	"github.com/manab-pr/nebulo/modules/transfers/domain/usecases"
	"github.com/manab-pr/nebulo/modules/transfers/presentation/http/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TransferHandler struct {
	getPendingUseCase *usecases.GetPendingTransfersUseCase
	completeUseCase   *usecases.CompleteTransferUseCase
	cancelUseCase     *usecases.CancelTransferUseCase
	validator         *validator.Validate
}

func NewTransferHandler(
	getPendingUseCase *usecases.GetPendingTransfersUseCase,
	completeUseCase *usecases.CompleteTransferUseCase,
	cancelUseCase *usecases.CancelTransferUseCase,
) *TransferHandler {
	return &TransferHandler{
		getPendingUseCase: getPendingUseCase,
		completeUseCase:   completeUseCase,
		cancelUseCase:     cancelUseCase,
		validator:         validator.New(),
	}
}

// GetPendingTransfers handles getting pending transfers for a device
func (h *TransferHandler) GetPendingTransfers(c *gin.Context) {
	deviceID := c.Param("deviceId")
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device ID is required"})
		return
	}

	transfers, err := h.getPendingUseCase.Execute(c.Request.Context(), deviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := dto.ToPendingTransferResponses(transfers)
	c.JSON(http.StatusOK, gin.H{
		"message": "Pending transfers retrieved successfully",
		"data":    responses,
	})
}

// CompleteTransfer handles transfer completion confirmation
func (h *TransferHandler) CompleteTransfer(c *gin.Context) {
	var req dto.CompleteTransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.completeUseCase.Execute(c.Request.Context(), req.ToEntity())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transfer completed successfully",
	})
}

// CancelTransfer handles transfer cancellation
func (h *TransferHandler) CancelTransfer(c *gin.Context) {
	transferID := c.Param("id")
	if transferID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transfer ID is required"})
		return
	}

	err := h.cancelUseCase.Execute(c.Request.Context(), transferID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transfer canceled successfully",
	})
}
