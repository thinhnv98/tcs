package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

// Giả lập một request body lớn với nhiều field
var largeJSON []byte

func init() {
	// Tạo một JSON lớn với nhiều field
	data := make(map[string]interface{})
	for i := 0; i < 30; i++ {
		key := fmt.Sprintf("%v%v", "field", i)
		data[key] = fmt.Sprintf("%v%v", "value", i)
	}
	// Thêm các field cần thiết
	data["field1"] = "value1"
	data["field2"] = 123

	largeJSON, _ = json.Marshal(data)
}

type RequestData struct {
	Field1 string `form:"field1" json:"field1"`
	Field2 int    `form:"field2" json:"field2"`
}

// Helper function to perform a request
func performRequest(r http.Handler, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
