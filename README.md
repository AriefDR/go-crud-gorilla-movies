# 🎬 Simple Go CRUD API with Gorilla Mux

This is a mini project that demonstrates a simple CRUD RESTful API using Go and the Gorilla Mux router.

## 🛠 Features

- ✅ Create, Read, Update, Delete (CRUD) operations on movie data
- 🧾 JSON API responses
- ❗ Error handling for:
  - Invalid request payloads (400)
  - Not found resources (404)
  - Internal server errors (500)
- 🌐 Custom 404 handler
- 📦 Vendoring enabled (via `vendor/`)

## 🧪 Sample API Endpoints

| Method | Endpoint           | Description          |
|--------|--------------------|----------------------|
| GET    | `/movies`          | Get all movies       |
| GET    | `/movies/{id}`     | Get movie by ID      |
| POST   | `/movies`          | Create a new movie   |
| PUT    | `/movies/{id}`     | Update a movie       |
| DELETE | `/movies/{id}`     | Delete a movie       |

## 🚀 Getting Started

1. Clone this repository:

```bash
git clone https://github.com/AriefDR/go-crud-gorilla-movies.git
```

2. Change to the project directory:
```bash
cd go-crud-gorilla-movies
```

3. (Optional) To ensure dependencies are correct, run:
```bash
go mod tidy
go mod vendor
```

4. Run the application using vendoring mode:
```bash
go run -mod=vendor main.go
```
