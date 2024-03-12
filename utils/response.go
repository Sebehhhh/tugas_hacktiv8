package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse adalah struktur untuk respons error JSON
type ErrorResponse struct {
	Message string `json:"message"`
}

// SendErrorResponse mengirim respons JSON dengan pesan error
func SendErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{Message: message})
}

// SendSuccessResponse mengirim respons JSON dengan data yang berhasil
func SendSuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

// SendNotFoundError mengirim respons JSON untuk kesalahan "tidak ditemukan"
func SendNotFoundError(c *gin.Context, message string) {
	SendErrorResponse(c, http.StatusNotFound, message)
}

// SendInternalServerError mengirim respons JSON untuk kesalahan "internal server error"
func SendInternalServerError(c *gin.Context, message string) {
	SendErrorResponse(c, http.StatusInternalServerError, message)
}

// SendBadRequestError mengirim respons JSON untuk kesalahan "permintaan tidak valid"
func SendBadRequestError(c *gin.Context, message string) {
	SendErrorResponse(c, http.StatusBadRequest, message)
}
