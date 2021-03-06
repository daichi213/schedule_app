// Package gofight offers simple API http handler testing for Golang framework.
//
// Details about the gofight project are found in github page:
//
//     https://github.com/appleboy/gofight
//
// Installation:
//
//    $ go get -u github.com/appleboy/gofight
//
// Set Header: You can add custom header via SetHeader func.
//
//    SetHeader(gofight.H{
//      "X-Version": version,
//    })
//
// Set Cookie: You can add custom cookie via SetCookie func.
//
//    SetCookie(gofight.H{
//      "foo": "bar",
//    })
//
// Set query string: Using SetQuery to generate query string data.
//
//    SetQuery(gofight.H{
//      "a": "1",
//      "b": "2",
//    })
//
// POST FORM Data: Using SetForm to generate form data.
//
//    SetForm(gofight.H{
//      "a": "1",
//      "b": "2",
//    })
//
// POST JSON Data: Using SetJSON to generate json data.
//
//    SetJSON(gofight.H{
//      "a": "1",
//      "b": "2",
//    })
//
// POST RAW Data: Using SetBody to generate raw data.
//
//    SetBody("a=1&b=1")
//
// For more details, see the documentation and example.
//
package gofight

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Media types
const (
	Version         = "1.0"
	UserAgent       = "User-Agent"
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
	ApplicationForm = "application/x-www-form-urlencoded"
)

// Timeout for http client
// see https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
var Timeout = time.Second * 10

// HTTPResponse is basic HTTP response type
type HTTPResponse *httptest.ResponseRecorder

// HTTPRequest is basic HTTP request type
type HTTPRequest *http.Request

// ResponseFunc response handling func type
type ResponseFunc func(HTTPResponse, HTTPRequest)

// H is HTTP Header Type
type H map[string]string

// D is HTTP Data Type
type D map[string]interface{}

// RequestConfig provide user input request structure
type RequestConfig struct {
	Method      string
	Path        string
	Body        string
	Headers     H
	Cookies     H
	Debug       bool
	ContentType string
}

// UploadFile for upload file struct
type UploadFile struct {
	Path    string
	Name    string
	Content []byte
}

// New supply initial structure
func New() *RequestConfig {

	return &RequestConfig{}
}

// SetDebug supply enable debug mode.
func (rc *RequestConfig) SetDebug(enable bool) *RequestConfig {
	rc.Debug = enable

	return rc
}

// GET is request method.
func (rc *RequestConfig) GET(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "GET"

	return rc
}

// POST is request method.
func (rc *RequestConfig) POST(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "POST"

	return rc
}

// PUT is request method.
func (rc *RequestConfig) PUT(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "PUT"

	return rc
}

// DELETE is request method.
func (rc *RequestConfig) DELETE(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "DELETE"

	return rc
}

// PATCH is request method.
func (rc *RequestConfig) PATCH(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "PATCH"

	return rc
}

// HEAD is request method.
func (rc *RequestConfig) HEAD(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "HEAD"

	return rc
}

// OPTIONS is request method.
func (rc *RequestConfig) OPTIONS(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "OPTIONS"

	return rc
}

// SetHeader supply http header what you defined.
func (rc *RequestConfig) SetHeader(headers H) *RequestConfig {
	if len(headers) > 0 {
		rc.Headers = headers
	}

	return rc
}

// SetJSON supply JSON body.
func (rc *RequestConfig) SetJSON(body D) *RequestConfig {
	if b, err := json.Marshal(body); err == nil {
		rc.Body = string(b)
	}

	return rc
}

// SetJSONInterface supply JSON body
func (rc *RequestConfig) SetJSONInterface(body interface{}) *RequestConfig {
	if b, err := json.Marshal(body); err == nil {
		rc.Body = string(b)
	}

	return rc
}

// SetForm supply form body.
func (rc *RequestConfig) SetForm(body H) *RequestConfig {
	f := make(url.Values)

	for k, v := range body {
		f.Set(k, v)
	}

	rc.Body = f.Encode()

	return rc
}

// SetFileFromPath upload new file.
func (rc *RequestConfig) SetFileFromPath(uploads []UploadFile, params ...H) *RequestConfig {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for _, f := range uploads {
		reader := bytes.NewReader(f.Content)
		if reader.Size() == 0 {
			file, err := os.Open(f.Path)
			if err != nil {
				log.Fatal(err)
			}

			defer file.Close()
			part, err := writer.CreateFormFile(f.Name, filepath.Base(f.Path))
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(part, file)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			part, err := writer.CreateFormFile(f.Name, filepath.Base(f.Path))
			if err != nil {
				log.Fatal(err)
			}
			_, err = reader.WriteTo(part)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if len(params) > 0 {
		for key, val := range params[0] {
			_ = writer.WriteField(key, val)
		}
	}

	err := writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	rc.ContentType = writer.FormDataContentType()
	rc.Body = body.String()

	return rc
}

// SetQuery supply query string.
func (rc *RequestConfig) SetQuery(query H) *RequestConfig {
	f := make(url.Values)

	for k, v := range query {
		f.Set(k, v)
	}

	if strings.Contains(rc.Path, "?") {
		rc.Path = rc.Path + "&" + f.Encode()
	} else {
		rc.Path = rc.Path + "?" + f.Encode()
	}

	return rc
}

// SetBody supply raw body.
func (rc *RequestConfig) SetBody(body string) *RequestConfig {
	if len(body) > 0 {
		rc.Body = body
	}

	return rc
}

// SetCookie supply cookies what you defined.
func (rc *RequestConfig) SetCookie(cookies H) *RequestConfig {
	if len(cookies) > 0 {
		rc.Cookies = cookies
	}

	return rc
}

func (rc *RequestConfig) initTest() (*http.Request, *httptest.ResponseRecorder) {
	qs := ""
	if strings.Contains(rc.Path, "?") {
		ss := strings.Split(rc.Path, "?")
		rc.Path = ss[0]
		qs = ss[1]
	}

	body := bytes.NewBufferString(rc.Body)

	req, _ := http.NewRequest(rc.Method, rc.Path, body)

	if len(qs) > 0 {
		req.URL.RawQuery = qs
	}

	// Auto add user agent
	req.Header.Set(UserAgent, "Gofight-client/"+Version)

	if rc.Method == "POST" || rc.Method == "PUT" || rc.Method == "PATCH" {
		if strings.HasPrefix(rc.Body, "{") {
			req.Header.Set(ContentType, ApplicationJSON)
		} else {
			req.Header.Set(ContentType, ApplicationForm)
		}
	}

	if rc.ContentType != "" {
		req.Header.Set(ContentType, rc.ContentType)
	}

	if len(rc.Headers) > 0 {
		for k, v := range rc.Headers {
			req.Header.Set(k, v)
		}
	}

	if len(rc.Cookies) > 0 {
		for k, v := range rc.Cookies {
			req.AddCookie(&http.Cookie{Name: k, Value: v})
		}
	}

	if rc.Debug {
		log.Printf("Request Method: %s", rc.Method)
		log.Printf("Request Path: %s", rc.Path)
		log.Printf("Request Body: %s", rc.Body)
		log.Printf("Request Headers: %s", rc.Headers)
		log.Printf("Request Cookies: %s", rc.Cookies)
		log.Printf("Request Header: %s", req.Header)
	}

	w := httptest.NewRecorder()

	return req, w
}

// Run execute http request
func (rc *RequestConfig) Run(r http.Handler, response ResponseFunc) {
	req, w := rc.initTest()
	r.ServeHTTP(w, req)
	response(w, req)
}
