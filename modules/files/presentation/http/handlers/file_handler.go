package handlers

import (
	"io"
	"net/http"

	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/files/domain/usecases"
	"github.com/manab-pr/nebulo/modules/files/presentation/http/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	maxUploadMemoryMB = 32
	mbShift           = 20 // 1 << 20 = 1MB
)

type FileHandler struct {
	storeUseCase  *usecases.StoreFileUseCase
	getUseCase    *usecases.GetFileUseCase
	deleteUseCase *usecases.DeleteFileUseCase
	validator     *validator.Validate
}

func NewFileHandler(
	storeUseCase *usecases.StoreFileUseCase,
	getUseCase *usecases.GetFileUseCase,
	deleteUseCase *usecases.DeleteFileUseCase,
) *FileHandler {
	return &FileHandler{
		storeUseCase:  storeUseCase,
		getUseCase:    getUseCase,
		deleteUseCase: deleteUseCase,
		validator:     validator.New(),
	}
}

// StoreFile handles file upload and storage
func (h *FileHandler) StoreFile(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Parse multipart form
	err := c.Request.ParseMultipartForm(maxUploadMemoryMB << mbShift) // 32 MB max memory
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	// Get file from form
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}
	defer file.Close()

	// Read file data
	fileData, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	// Create store request
	req := dto.StoreFileRequest{
		Name:         header.Filename,
		Size:         int64(len(fileData)),
		MimeType:     header.Header.Get("Content-Type"),
		TargetDevice: c.PostForm("target_device"),
	}

	if validationErr := h.validator.Struct(req); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	storedFile, err := h.storeUseCase.Execute(c.Request.Context(), userID, req.ToEntity(), fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.ToFileResponse(storedFile)
	c.JSON(http.StatusCreated, gin.H{
		"message": "File stored successfully",
		"data":    response,
	})
}

// GetFile handles file metadata retrieval
func (h *FileHandler) GetFile(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	fileID := c.Param("fileId")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	file, err := h.getUseCase.Execute(c.Request.Context(), userID, fileID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := dto.ToFileResponse(file)
	c.JSON(http.StatusOK, gin.H{
		"message": "File retrieved successfully",
		"data":    response,
	})
}

// GetAllFiles handles listing all files
func (h *FileHandler) GetAllFiles(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	files, err := h.getUseCase.GetAllFiles(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := dto.ToFileResponses(files)
	c.JSON(http.StatusOK, gin.H{
		"message": "Files retrieved successfully",
		"data":    responses,
	})
}

// DeleteFile handles file deletion
func (h *FileHandler) DeleteFile(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	fileID := c.Param("fileId")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	err := h.deleteUseCase.Execute(c.Request.Context(), userID, fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File deleted successfully",
	})
}
