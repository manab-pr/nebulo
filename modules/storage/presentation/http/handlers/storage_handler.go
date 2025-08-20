package handlers

import (
	"net/http"

	"github.com/manab-pr/nebulo/modules/storage/domain/usecases"

	"github.com/gin-gonic/gin"
)

type StorageHandler struct {
	summaryUseCase       *usecases.GetStorageSummaryUseCase
	deviceStorageUseCase *usecases.GetDeviceStorageUseCase
}

func NewStorageHandler(
	summaryUseCase *usecases.GetStorageSummaryUseCase,
	deviceStorageUseCase *usecases.GetDeviceStorageUseCase,
) *StorageHandler {
	return &StorageHandler{
		summaryUseCase:       summaryUseCase,
		deviceStorageUseCase: deviceStorageUseCase,
	}
}

// GetStorageSummary handles getting storage summary across all devices
func (h *StorageHandler) GetStorageSummary(c *gin.Context) {
	summary, err := h.summaryUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Storage summary retrieved successfully",
		"data":    summary,
	})
}

// GetDeviceStorage handles getting storage info for a specific device
func (h *StorageHandler) GetDeviceStorage(c *gin.Context) {
	deviceID := c.Param("deviceId")
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device ID is required"})
		return
	}

	storageInfo, err := h.deviceStorageUseCase.Execute(c.Request.Context(), deviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Device storage info retrieved successfully",
		"data":    storageInfo,
	})
}
