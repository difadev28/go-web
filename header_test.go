package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeaderCustom(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Difa Ardiansyah")
	fmt.Fprint(writer, "OK")
}
func TestResponseHeaderCustom(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	ResponseHeaderCustom(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Header.Get("X-Powered-By"))
}
