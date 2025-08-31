# Go CRUD API

A simple **CRUD REST API** built in **Go** using a **clean folder architecture**.  
This project demonstrates a layered approach with **handlers, services, and generic in-memory storage**.  

---

## 📂 Project Structure
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


---

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


