package main

// import (
// 	"io"
// 	"bytes"
// 	"testing"
// 	"net/http/httptest"
// 	"github.com/stretchr/testify/assert"
// )

// func testRoothandler(t *testing.T){
// 	router := RootRouter()

// 	req := httptest.NewRequest("GET", "/", nil)
// 	rec := httptest.NewRecorder()

// 	router.serveHTTP(rec, req)

// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, http.Status)
// }

// func testPosthandler(t *testing.T){
// 	router := PostRouter()
// 	reqbody := strings.NewReader(`{text: "hello world", status: false}`)
// 	req := httptest.NewRequest("POST", "/post", reqbody)

// 	rec := httptest.NewRecorder()

// 	router.serveHTTP(rec, req)

// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	// assert.Equal(t, http.Status)
// }

// func testShowhandler(t *testing.T){
// 	router := PostRouter()

// 	req := httptest.NewRequest("GET", "/1", nil)
// 	rec := httptest.NewRecorder()

// 	router.serveHTTP(rec, req)

// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, http.Status)
// }

// func testUpdatehandler(t *testing.T){
// 	router := PostRouter()

// 	req := httptest.NewRequest("POST", "/1/update", {text: "hello world.", status: true})
// 	rec := httptest.NewRecorder()

// 	router.serveHTTP(rec, req)

// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, http.Status)
// }