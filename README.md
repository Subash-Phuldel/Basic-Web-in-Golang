# Basic-Web-in-Golang
# Basic-Web-in-Golang
# Basic-Web-in-Golang


# KnowledgeHub – REST API Development in Go

KnowledgeHub is a simple REST API built using Go’s standard `net/http` package.
It allows users to create, retrieve, and delete articles.
The project also demonstrates:

* REST API design
* JSON handling
* URL routing with path parameters
* Slug generation
* Template rendering
* Static file serving
* Concurrency safety using mutex
* Unique ID generation using Snowflake

---

# 🚀 Features

* Create new articles
* Retrieve article by slug
* Delete article by slug
* Server-side HTML rendering using Go templates
* Static file serving
* Thread-safe in-memory storage
* Unique ID generation using Snowflake algorithm

---

# 📂 Project Structure

```
knowledgehub/
│
├── cmd/web/
│   ├── main.go
│   ├── handlers.go
│   ├── models.go
│   ├── validation.go
│   └── slug.go
│
├── ui/
│   ├── html/pages/
│   │   ├── base.html
│   │   ├── nav.html
│   │   ├── main.html
│   │   └── footer.html
│   │
│   └── static/
│
└── README.md
```

---

# 🧠 Core Concepts Used

## 1. In-Memory Storage

Articles are stored in a slice:

```go
var articles = make([]Article, 0, 20)
```

Since multiple HTTP requests can run at the same time, a `sync.Mutex` is used:

```go
var mu sync.Mutex
```

This ensures safe concurrent access.

---

## 2. Unique ID Generation (Snowflake)

The project uses:

```
github.com/bwmarrin/snowflake
```

Each article gets a globally unique ID:

```go
id := node.Generate().String()
```

This is useful in distributed systems.

---

## 3. Slug Generation

The `createSlug()` function:

* Converts title to lowercase
* Removes special characters
* Replaces spaces with `-`
* Trims extra dashes

Example:

```
"Hello World!! Go Lang"
→ "hello-world-go-lang"
```

---

## 4. Validation

Minimum length validation is done using:

```go
func minLength(text string, minLen int) error
```

Used to validate:

* Article title (minimum 5 characters)
* Article body (minimum 10 characters)
* Slug length

---

## 5. Routing (Go 1.22+ Pattern Routing)

The project uses the new HTTP routing syntax:

```go
mux.HandleFunc("POST /articles", postArticleHandler)
mux.HandleFunc("GET /articles/{slug}", getArticleHandler)
mux.HandleFunc("DELETE /articles/{slug}", deleteArticleHandler)
```

Path parameter extraction:

```go
slug := r.PathValue("slug")
```

---

# 🌐 API Endpoints

---

## 1️⃣ Home (Text Response)

### GET `/`

Response:

```
Welcome to KnowledgeHub
```

---

## 2️⃣ Render HTML Home Page

### GET `/home`

Renders HTML templates:

* base.html
* nav.html
* main.html
* footer.html

Template composition is done using:

```go
ExecuteTemplate(w, "base", nil)
```

---

## 3️⃣ Create Article

### POST `/articles`

### Request Body (JSON)

```json
{
  "title": "My First Article",
  "body": "This is the body of my article."
}
```

### Success Response (201 Created)

```json
{
  "id": "1700000000000",
  "slug": "my-first-article",
  "title": "My First Article",
  "body": "This is the body of my article."
}
```

### Validations

* Title must be at least 5 characters
* Body must be at least 10 characters

---

## 4️⃣ Get Article by Slug

### GET `/articles/{slug}`

Example:

```
GET /articles/my-first-article
```

### Success Response

```json
{
  "id": "...",
  "slug": "my-first-article",
  "title": "My First Article",
  "body": "This is the body of my article."
}
```

### Errors

* 400 → slug too short
* 404 → article not found

---

## 5️⃣ Delete Article

### DELETE `/articles/{slug}`

Example:

```
DELETE /articles/my-first-article
```

### Success Response

```
204 No Content
content removed
```

---

# 📦 Static File Serving

Static files are served using:

```go
fileServer := http.FileServer(http.Dir("../../ui/static"))
mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
```

Access static files via:

```
http://localhost:8080/static/filename.css
```

---

# 🖥 How to Run

### 1️⃣ Install Dependencies

```bash
go get github.com/bwmarrin/snowflake
```

### 2️⃣ Run Server

```bash
go run ./cmd/web
```

Server starts at:

```
http://localhost:8080
```

---

# 🛠 Technologies Used

* Go (net/http)
* Go Templates
* Snowflake ID generator
* Mutex for concurrency safety
* JSON encoding/decoding

---

# ⚠ Limitations

* Data is stored in memory (not persistent)
* Server restart clears all articles
* No database integration
* No authentication
* No pagination

---

# 📚 Learning Outcomes

This project demonstrates:

* Building REST APIs without frameworks
* Working with JSON
* Handling HTTP methods properly
* Using route parameters
* Template composition
* Static file serving
* Writing thread-safe code
* Structuring a Go web application

---

# 🔮 Possible Improvements

* Add PUT endpoint for updating articles
* Add database (PostgreSQL)
* Add pagination
* Add middleware (logging, recovery)
* Add authentication (JWT)
* Add unit tests
* Convert to clean architecture

---
