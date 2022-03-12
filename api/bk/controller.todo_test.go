package main

import (
	"bytes"
	// "strings"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func TestShowTodoIndex(t *testing.T) {
	router = gin.Default()

	initializeRoutes()

	req := httptest.NewRequest("GET", "/", bytes.NewBufferString("Empty Message"))

	testHttpResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		return statusOK
	})
}

func TestShowTodoContent(t *testing.T) {
	router = gin.Default()

	initializeRoutes()

	req := httptest.NewRequest("GET", "/7", bytes.NewBufferString("Empty Message"))

	testHttpResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		return statusOK
	})
}

func TestShowTodoContentNotExist(t *testing.T) {
	router = gin.Default()

	initializeRoutes()

	req := httptest.NewRequest("GET", "/9999", bytes.NewBufferString("Empty Message"))

	testHttpResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusNotFound
		return statusOK
	})
}


// ControllerとModelをまとめてテストするタイプ
// Controller単体のテストはないが、それ以外の関数のテストを用意するようにする
func TestCreateTodo(t *testing.T) {

	nowTime := GetTimeNumber()

	Title := nowTime
	Content := "This messages is test message in test1."
	Status := "0"

	router = gin.Default()
	initializeRoutes()

	req := httptest.NewRequest("POST", "/new", bytes.NewBufferString("Empty Message"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	params := req.URL.Query()
	params.Add("Title", Title)
	params.Add("Content", Content)
	params.Add("Status", Status)
	req.URL.RawQuery = params.Encode()

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	db = GetDB()
	DB, err := db.DB()
	defer DB.Close()
	if err != nil {
		panic(err)
	}
	db.Debug().Last(&todo)

	assert.Equal(t, nowTime, todo.Title)
}