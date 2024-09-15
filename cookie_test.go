package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Token"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success create cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-Token")
	if err != nil {
		panic(err)
	} else {
		fmt.Fprintf(writer, "Ini cookie %s", cookie)
	}

}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	serve := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080?name=Difa", nil)
	recorder := httptest.NewRecorder()
	SetCookie(recorder, request)
	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Coookie %s \n Value : %s", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080?name=Difa", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-Token"
	cookie.Value = "DIfa"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()
	GetCookie(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
