package storage

type MemoryStore[T any] struct {
	Items  []T
	NextID int
}

func NewMemoryStore[T any]() *MemoryStore[T] {
	return &MemoryStore[T]{
		Items:  []T{},
		NextID: 1,
	}
}

// Add an item
func (m *MemoryStore[T]) Add(item T) T {
	m.Items = append(m.Items, item)
	m.NextID++
	return item
}

// Get all items
func (m *MemoryStore[T]) GetAll() []T {
	return m.Items
}

// Replace all items (used for update/delete)
func (m *MemoryStore[T]) SetAll(items []T) {
	m.Items = items
}
