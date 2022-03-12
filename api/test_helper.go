package main

import (
	"time"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
)

// もし、ginのみでMVCのアプリとして開発する場合はtemplatesディレクトリ以下にhtmlファイルを格納して出力タグのテストも行う
func getRouter(withTemplate bool) *gin.Engine {
	r := gin.Default()
	if withTemplate {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func testHttpResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func SetupFindFuncs() {
	user := []User{
		{Name: "user1", Email: "testuser1@example.com", Password: "testpassword"},
	}

	db.Model(&user).Create(&user)

	testTodos := []Todo {
		{Title: "test1", Content:"This message is in test1.", Status: 1},
		{Title: "test2", Content:"This message is in test2.", Status: 1},
		{Title: "test3", Content:"This message is in test3.", Status: 1},
	}

	db.Model(&todo).Create(&testTodos)
}

func GetTimeNumber() string {
	// 以下の時刻をもじって時刻のフォーマットを指定する
	// layout = "Now, Monday Jan 02 15:04:05 JST 2006"
	layout := "20060102150405"
	nowTime := time.Now()
	return nowTime.Format(layout)
}