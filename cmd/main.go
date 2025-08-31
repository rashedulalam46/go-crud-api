package main

import (
	"fmt"
	"gocrudapi/internal/storage"
	"gocrudapi/internal/user"
	"net/http"
)

func main() {
	store := storage.NewMemoryStore[user.User]()
	service := user.NewService(store)
	handler := user.NewHandler(service)

	http.HandleFunc("/create", handler.CreateUser)
	http.HandleFunc("/users", handler.GetUsers)
	http.HandleFunc("/user", handler.GetUser)
	http.HandleFunc("/update", handler.UpdateUser)
	http.HandleFunc("/delete", handler.DeleteUser)

	fmt.Println("🚀 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
