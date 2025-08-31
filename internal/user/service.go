package user

import "gocrudapi/internal/storage"

type Service struct {
	store *storage.MemoryStore[User]
}

func NewService(store *storage.MemoryStore[User]) *Service {
	return &Service{store: store}
}

func (s *Service) Create(name string) User {
	u := User{ID: s.store.NextID, Name: name}
	items := s.store.GetAll()
	items = append(items, u)
	s.store.SetAll(items)
	s.store.NextID++
	return u
}

func (s *Service) GetAll() []User {
	return s.store.GetAll()
}

func (s *Service) GetByID(id int) (User, bool) {
	for _, u := range s.store.GetAll() {
		if u.ID == id {
			return u, true
		}
	}
	return User{}, false
}

func (s *Service) Update(id int, name string) (User, bool) {
	items := s.store.GetAll()
	for i, u := range items {
		if u.ID == id {
			items[i].Name = name
			s.store.SetAll(items)
			return items[i], true
		}
	}
	return User{}, false
}

func (s *Service) Delete(id int) bool {
	items := s.store.GetAll()
	for i, u := range items {
		if u.ID == id {
			items = append(items[:i], items[i+1:]...)
			s.store.SetAll(items)
			return true
		}
	}
	return false
}
