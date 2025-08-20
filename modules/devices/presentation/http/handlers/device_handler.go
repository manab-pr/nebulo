package handlers

import (
	"net/http"

	"github.com/manab-pr/nebulo/modules/devices/domain/usecases"
	"github.com/manab-pr/nebulo/modules/devices/presentation/http/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DeviceHandler struct {
	registerUseCase    *usecases.RegisterDeviceUseCase
	heartbeatUseCase   *usecases.HeartbeatUseCase
	listDevicesUseCase *usecases.ListDevicesUseCase
	deleteUseCase      *usecases.DeleteDeviceUseCase
	validator          *validator.Validate
}

func NewDeviceHandler(
	registerUseCase *usecases.RegisterDeviceUseCase,
	heartbeatUseCase *usecases.HeartbeatUseCase,
	listDevicesUseCase *usecases.ListDevicesUseCase,
	deleteUseCase *usecases.DeleteDeviceUseCase,
) *DeviceHandler {
	return &DeviceHandler{
		registerUseCase:    registerUseCase,
		heartbeatUseCase:   heartbeatUseCase,
		listDevicesUseCase: listDevicesUseCase,
		deleteUseCase:      deleteUseCase,
		validator:          validator.New(),
	}
}

// RegisterDevice handles device registration
func (h *DeviceHandler) RegisterDevice(c *gin.Context) {
	var req dto.DeviceRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	device, err := h.registerUseCase.Execute(c.Request.Context(), req.ToEntity())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.ToDeviceResponse(device)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Device registered successfully",
		"data":    response,
	})
}

// Heartbeat handles device heartbeat
func (h *DeviceHandler) Heartbeat(c *gin.Context) {
	var req dto.DeviceHeartbeatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.heartbeatUseCase.Execute(c.Request.Context(), req.ToEntity())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Heartbeat received successfully",
	})
}

// GetDevices handles listing all devices
func (h *DeviceHandler) GetDevices(c *gin.Context) {
	devices, err := h.listDevicesUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := dto.ToDeviceResponses(devices)
	c.JSON(http.StatusOK, gin.H{
		"message": "Devices retrieved successfully",
		"data":    responses,
	})
}

// DeleteDevice handles device removal
func (h *DeviceHandler) DeleteDevice(c *gin.Context) {
	deviceID := c.Param("id")
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device ID is required"})
		return
	}

	err := h.deleteUseCase.Execute(c.Request.Context(), deviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Device deleted successfully",
	})
}