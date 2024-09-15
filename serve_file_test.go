package go_web

import (
	_ "embed" // {{ edit_1 }}
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query().Get("name")
	if query != "" {
		http.ServeFile(writer, request, "./resources/not_empty.html")
	} else {
		http.ServeFile(writer, request, "./resources/empty.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/not_empty.html
var resourceOK string

//go:embed resources/empty.html
var resourceNotOK string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query().Get("name")
	if query != "" {
		fmt.Fprint(writer, resourceOK)
	} else {
		fmt.Fprint(writer, resourceNotOK)
	}
}
func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
