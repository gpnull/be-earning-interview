package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRecordAttendanceAPI(t *testing.T) {
	// Thiết lập Gin router
	router := gin.Default()
	router.POST("/api/attendance", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Attendance recorded"})
	})

	// Tạo request và recorder
	req, _ := http.NewRequest("POST", "/api/attendance", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Kiểm tra kết quả
	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "Attendance recorded")
}

func TestGetAttendanceByEmployeeAPI(t *testing.T) {
	// Thiết lập Gin router
	router := gin.Default()
	router.GET("/api/attendance/:employeeId", func(c *gin.Context) {
		employeeId := c.Param("employeeId")
		c.JSON(http.StatusOK, gin.H{"employeeId": employeeId, "records": "Some records"})
	})

	// Tạo request và recorder
	req, _ := http.NewRequest("GET", "/api/attendance/123", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Kiểm tra kết quả
	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "Some records")
}

func TestGetAttendanceByEmployeeAPI_InvalidID(t *testing.T) {
	router := gin.Default()
	router.GET("/api/attendance/:employeeId", func(c *gin.Context) {
		employeeId := c.Param("employeeId")
		if employeeId != "123" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"employeeId": employeeId, "records": "Some records"})
	})

	req, _ := http.NewRequest("GET", "/api/attendance/abc", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid employee ID")
}
