package user

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	created := h.service.Create(u.Name)
	json.NewEncoder(w).Encode(created)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.service.GetAll())
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	u, found := h.service.GetByID(id)
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(u)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, ok := h.service.Update(id, u.Name)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	if !h.service.Delete(id) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("User deleted"))
}
