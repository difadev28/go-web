package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Difa",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
func TemplateFunctionLen(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Difa",
	})
}
func TestTemplateFunctionLen(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionLen(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{upper .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Difa",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionPipelines(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{.SayHello "Budi" | upper}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Difa",
	})
}
func TestTemplateFunctionPipelines(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipelines(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
