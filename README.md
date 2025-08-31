# Go CRUD API

A simple **CRUD REST API** built in **Go** using a **clean folder architecture**.  
This project demonstrates a layered approach with **handlers, services, and generic in-memory storage**.  

---

## 📂 Project Structure
```
GoCrudApi/
├── cmd/
│ └── main.go # Entry point
├── internal/
│ ├── storage/
│ │ └── memory.go # Generic in-memory storage
│ └── user/
│ ├── handler.go # HTTP handlers
│ ├── model.go # User model
│ └── service.go # Business logic
├── go.mod
└── go.sum

```
## ⚙️ Setup

### 1. Clone the project
```bash
git clone https://github.com/rashedulalam46/go-crud-api.git
cd GoCrudApi
```

### 2. Initialize dependencies
```bash
go mod tidy
```

### 3. Run the API
```bash
go run ./cmd
```

Server will start on: http://localhost:8080

## 🚀 API Endpoints

| Action        | Method | URL               | Body (JSON)                   |
| ------------- | ------ | ----------------- | ----------------------------- |
| Create User   | POST   | `/create`         | `{ "name": "Alice" }`         |
| Get All Users | GET    | `/users`          | None                          |
| Get User      | GET    | `/user?id=<id>`   | None                          |
| Update User   | PUT    | `/update?id=<id>` | `{ "name": "Alice Updated" }` |
| Delete User   | DELETE | `/delete?id=<id>` | None                          |



## 🛠 Features
- Clean folder architecture (cmd, internal/service, internal/handler, internal/storage)
- Generic in-memory storage using Go generics
- Full CRUD operations (Create, Read, Update, Delete)
- Ready to extend with real database



