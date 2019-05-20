package storage

import (
	"go-playground/model"
	"time"
)

// MemoryDeskStorage stores all desks in memory only
type MemoryDeskStorage struct {
	desks map[string]model.Desk
}

// NewDesk adds new desk to repo
func (storage MemoryDeskStorage) NewDesk(desk model.Desk) bool {
	storage.desks[desk.Name] = desk
	return true
}

// GetDesk returns desk by name
func (storage MemoryDeskStorage) GetDesk(name string) (model.Desk, bool) {
	value, ok := storage.desks[name]
	return value, ok
}

// DeleteDesk deletes desk by name
func (storage MemoryDeskStorage) DeleteDesk(name string) bool {
	_, ok := storage.desks[name]
	delete(storage.desks, name)
	return ok
}

// GetAllDesks return all desks
func (storage MemoryDeskStorage) GetAllDesks() []model.Desk {
	copy := make([]model.Desk, 0, len(storage.desks))
	for _, desk := range storage.desks {
		copy = append(copy, desk)
	}
	return copy
}

// CreateMemoryDeskStorage creates predefined storage
func CreateMemoryDeskStorage() MemoryDeskStorage {
	storage := MemoryDeskStorage{desks: make(map[string]model.Desk)}

	storage.NewDesk(model.Desk{
		Name:      "test",
		Cards:     []string{"A", "B", "C"},
		State:     model.VotingState,
		Timestamp: time.Now(),
		Users: []model.DeskUser{
			model.DeskUser{Name: "user1"},
			model.DeskUser{Name: "user2"},
		}})

	return storage
}
