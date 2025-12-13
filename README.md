# 📚 Go Web Development Learning Repository

> Repository pembelajaran pengembangan web backend menggunakan Go language dengan pendekatan test-driven.

## 🎯 Tujuan Pembelajaran

Repository ini dirancang sebagai media pembelajaran komprehensif untuk menguasai pengembangan web backend menggunakan Go. Setiap konsep dipraktikkan melalui test case yang dapat dieksekusi, sehingga memudahkan pemahaman melalui eksperimen langsung.

## 📋 Daftar Materi

### 1. **Fundamental HTTP Server**
- Membuat HTTP server dengan `http.Server`
- Menggunakan `http.ServeMux` untuk routing
- Memahami HTTP methods dan request URI

### 2. **HTTP Handler & Routing**
- Request handler dengan `http.HandlerFunc`
- Custom handler types
- Parameter parsing dari URL

### 3. **Form Processing**
- Parsing form data dengan `request.ParseForm()`
- Handling POST requests
- Validasi form input

### 4. **Query Parameters**
- Extract query parameters dari URL
- Multiple values handling
- Query string manipulation

### 5. **HTTP Cookies**
- Setting dan reading cookies
- Cookie path configuration
- Session management basics

### 6. **Template Engine**
- Basic template parsing dan execution
- File-based templates (`template.ParseFiles`)
- Directory-based templates (`template.ParseGlob`)
- Embedded templates dengan `//go:embed`
- Template layouts dan partials
- Template actions, conditionals, dan loops
- Template functions dan pipelines
- Template caching
- Template auto-escaping untuk XSS prevention

### 7. **File Operations**
- File upload dengan multipart forms
- File serving dan downloads
- Static file handler
- File server configuration

### 8. **HTTP Headers**
- Setting custom headers
- Content-Type handling
- Header manipulation

### 9. **HTTP Redirects**
- Implementasi redirects
- Status code handling
- Permanent vs temporary redirects

### 10. **Response Codes**
- Setting appropriate HTTP status codes
- Custom response handling
- Error responses

### 11. **Middleware Pattern**
- Membuat custom middleware
- Middleware chaining
- Error handling middleware
- Recovery middleware
- Request/response interception

### 12. **Security - XSS Prevention**
- Template auto-escaping
- Manual HTML escaping
- XSS vulnerability demonstration
- Security best practices

### 13. **Static File Serving**
- Serving static files
- Using `http.FileServer`
- Path stripping dengan `http.StripPrefix`

### 14. **Session Management** *(Coming Soon)*
- Session middleware implementation
- Session-based authentication
- Session storage strategies

## 🗂️ Struktur Project

```
go-web/
├── templates/          # HTML template files
│   ├── layout/        # Template layouts
│   └── *.html         # Template pages
├── resources/         # Static assets
│   ├── css/          # Stylesheets
│   ├── js/           # JavaScript files
│   └── images/       # Image files
├── *_test.go         # Test files untuk setiap materi
├── main.go           # Main application
├── go.mod            # Go module definition
└── README.md         # Dokumentasi (ini)
```

## 🚀 Cara Menggunakan Repository

### Prerequisites
- Go 1.20 atau lebih tinggi
- Text editor atau IDE (VS Code, GoLand, dll)

### Menjalankan Test

Untuk menjalankan semua test:
```bash
go test -v ./...
```

Untuk menjalankan test spesifik:
```bash
# Test untuk HTTP server
go test -v -run TestServer

# Test untuk templates
go test -v -run TestTemplate

# Test untuk file upload
go test -v -run TestUploadFile
```

### Membaca Kode

Setiap file test mengandung:
1. **Setup** - Inisialisasi server atau handler
2. **Test Cases** - Berbagai skenario pengujian
3. **Examples** - Contoh implementasi praktis
4. **Comments** - Penjelasan konsep dan best practices

## 📖 Metode Pembelajaran

### 1. **Test-Driven Learning**
Setiap konsep dipelajari melalui test case yang sudah disediakan. Anda dapat:
- Membaca kode untuk memahami konsep
- Menjalankan test untuk melihat hasilnya
- Memodifikasi test untuk eksperimen

### 2. **Progressive Complexity**
Materi disusun dari yang paling dasar hingga advanced:
- Server dasar → Routing → Templates → Middleware → Security

### 3. **Hands-on Practice**
Setiap materi memiliki test yang dapat dieksekusi langsung, memungkinkan:
- Eksperimen dengan code
- Melihat hasil instantly
- Debug dan troubleshooting

## 🛠️ Contoh Implementasi

### Basic HTTP Server
```go
func TestServer(t *testing.T) {
    server := http.Server{
        Addr: "localhost:8080",
        Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprint(w, "Hello World")
        }),
    }

    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
```

### Template Rendering
```go
func TestTemplate(t *testing.T) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := map[string]interface{}{
            "Title": "Welcome",
            "Name":  "Go Developer",
        }
        tmpl.Execute(w, data)
    })
}
```

### Middleware
```go
func TestMiddleware(t *testing.T) {
    middleware := func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Pre-processing
            t.Log("Request received")

            // Call next handler
            next.ServeHTTP(w, r)

            // Post-processing
            t.Log("Request completed")
        })
    }
}
```

## 🔧 Development Tips

### Best Practices
1. **Selalu buat test** untuk setiap fitur baru
2. **Gunakan middleware** untuk cross-cutting concerns
3. **Validate input** untuk security
4. **Handle errors** secara graceful
5. **Use templates** dengan auto-escaping untuk XSS prevention

### Testing Tips
- Gunakan `httptest.NewRecorder()` untuk testing handlers
- Mock external dependencies
- Test untuk error scenarios, bukan hanya success cases

## 📚 Referensi Tambahan

- [Official Go Web Programming Tutorial](https://golang.org/doc/articles/wiki/)
- [Go by Example: HTTP Servers](https://gobyexample.com/http-servers)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Net/http Package Documentation](https://golang.org/pkg/net/http/)

## 🤝 Kontribusi

Repository ini dibuat untuk pembelajaran. Jika Anda menemukan:
- **Bug** → Buat issue dengan detail error
- **Improvement** → Buat PR dengan perubahan
- **Question** → Tanyakan di issues

## 📄 License

Repository ini open-source untuk tujuan pembelajaran. Silakan gunakan, modifikasi, dan bagikan sesuai kebutuhan.

---

💡 **Happy Learning!** Semoga repository ini membantu Anda menguasai Go web development! 🚀