package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type InternalDeviceHandler struct {
	storagePath string
}

func NewInternalDeviceHandler(storagePath string) *InternalDeviceHandler {
	return &InternalDeviceHandler{
		storagePath: storagePath,
	}
}

// StoreFile handles incoming file storage requests from backend
func (h *InternalDeviceHandler) StoreFile(c *gin.Context) {
	// Parse multipart form
	err := c.Request.ParseMultipartForm(32 << 20) // 32 MB max memory
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

	// Create storage directory if it doesn't exist
	err = os.MkdirAll(h.storagePath, 0755)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create storage directory"})
		return
	}

	// Create destination file
	fileName := c.PostForm("filename")
	if fileName == "" {
		fileName = header.Filename
	}
	
	dst, err := os.Create(filepath.Join(h.storagePath, fileName))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
		return
	}
	defer dst.Close()

	// Copy file contents
	_, err = io.Copy(dst, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "File stored successfully",
		"filename": fileName,
	})
}

// GetFile serves a file stored on this device
func (h *InternalDeviceHandler) GetFile(c *gin.Context) {
	fileID := c.Param("id")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	// Find file by ID (in practice, you'd have a mapping of ID to filename)
	// For now, we'll use the ID as filename
	filePath := filepath.Join(h.storagePath, fileID)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.File(filePath)
}

// GetStorageInfo reports current available storage
func (h *InternalDeviceHandler) GetStorageInfo(c *gin.Context) {
	// Get storage info for the storage path
	// This is a simplified version - in practice you'd want to get disk usage stats
	totalSpace := int64(100 * 1024 * 1024 * 1024) // 100GB default
	usedSpace := int64(0)

	// Walk through storage directory to calculate used space
	err := filepath.Walk(h.storagePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip files with errors
		}
		if !info.IsDir() {
			usedSpace += info.Size()
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate storage usage"})
		return
	}

	availableSpace := totalSpace - usedSpace

	c.JSON(http.StatusOK, gin.H{
		"total_storage":     totalSpace,
		"used_storage":      usedSpace,
		"available_storage": availableSpace,
	})
}

// ConfirmFile confirms file successfully received and stored
func (h *InternalDeviceHandler) ConfirmFile(c *gin.Context) {
	fileID := c.Param("fileId")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	// Check if file exists
	filePath := filepath.Join(h.storagePath, fileID)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File confirmed successfully",
		"file_id": fileID,
	})
}