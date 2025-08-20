package handlers

import (
	"net/http"

	"github.com/manab-pr/nebulo/modules/files/presentation/http/dto"
	"github.com/manab-pr/nebulo/modules/search/domain/usecases"

	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	searchFilesUseCase    *usecases.SearchFilesUseCase
	getLocationUseCase    *usecases.GetFileLocationUseCase
}

func NewSearchHandler(
	searchFilesUseCase *usecases.SearchFilesUseCase,
	getLocationUseCase *usecases.GetFileLocationUseCase,
) *SearchHandler {
	return &SearchHandler{
		searchFilesUseCase: searchFilesUseCase,
		getLocationUseCase: getLocationUseCase,
	}
}

// SearchFiles handles file search across all devices
func (h *SearchHandler) SearchFiles(c *gin.Context) {
	query := c.Query("name")

	files, err := h.searchFilesUseCase.Execute(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := dto.ToFileResponses(files)
	c.JSON(http.StatusOK, gin.H{
		"message": "Files search completed successfully",
		"data":    responses,
		"query":   query,
	})
}

// GetFileLocation handles getting file location information
func (h *SearchHandler) GetFileLocation(c *gin.Context) {
	fileID := c.Param("fileId")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	location, err := h.getLocationUseCase.Execute(c.Request.Context(), fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File location retrieved successfully",
		"data":    location,
	})
}