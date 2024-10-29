package main

import (
	"bytes"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"testing"
)

func BenchmarkShouldBindJSON(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.POST("/", func(c *gin.Context) {
		var data RequestData
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusOK)
	})

	req, _ := http.NewRequest("POST", "/", bytes.NewReader(largeJSON))
	req.Header.Set("Content-Type", "application/json")

	for i := 0; i < b.N; i++ {
		_ = performRequest(router, req)
	}
}

func BenchmarkShouldBind(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.POST("/", func(c *gin.Context) {
		var data RequestData
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusOK)
	})

	req, _ := http.NewRequest("POST", "/", bytes.NewReader(largeJSON))
	req.Header.Set("Content-Type", "application/json")

	for i := 0; i < b.N; i++ {
		_ = performRequest(router, req)
	}
}

func BenchmarkShouldBindQuery(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		var data RequestData
		if err := c.ShouldBindQuery(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusOK)
	})

	params := url.Values{}
	params.Add("field1", "value1")
	params.Add("field2", "123")
	req, _ := http.NewRequest("GET", "/?"+params.Encode(), nil)

	for i := 0; i < b.N; i++ {
		w := performRequest(router, req)
		if w.Code != http.StatusOK {
			b.Fatalf("Expected status 200 but got %d", w.Code)
		}
	}
}

func BenchmarkJSONParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		field1, err := jsonparser.GetString(largeJSON, "field1")
		if err != nil {
			b.Fatal(err)
		}
		field2, err := jsonparser.GetInt(largeJSON, "field2")
		if err != nil {
			b.Fatal(err)
		}

		// Sử dụng field1 và field2 nếu cần
		_ = field1
		_ = field2
	}
}
