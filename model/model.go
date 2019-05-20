package model

import (
	"time"
)

// Desk struct
type Desk struct {
	Name      string     `json:"name"`
	Cards     []string   `json:"cards"`
	State     int        `json:"state"`
	Timestamp time.Time  `json:"timestamp"`
	Users     []DeskUser `json:"users"`
}

// DeskUser struct
type DeskUser struct {
	Name string `json:"name"`
	Card string `json:"card"`
}

const (
	// VotingState value
	VotingState = iota

	// DisplayState value
	DisplayState = iota
)

// NewDesk adds new desk to repo
func NewDesk(desk Desk) {
	desks[desk.Name] = desk
}

// GetDesk returns desk by name
func GetDesk(name string) (Desk, bool) {
	value, ok := desks[name]
	return value, ok
}

// DeleteDesk deletes desk by name
func DeleteDesk(name string) {
	delete(desks, name)
}

// GetAllDesks return all desks
func GetAllDesks() []Desk {
	copy := make([]Desk, 0, len(desks))
	for _, desk := range desks {
		copy = append(copy, desk)
	}
	return copy
}

var desks map[string]Desk

func init() {
	desks = make(map[string]Desk)

	NewDesk(Desk{
		Name:      "test",
		Cards:     []string{"A", "B", "C"},
		State:     VotingState,
		Timestamp: time.Now(),
		Users: []DeskUser{
			DeskUser{Name: "user1"},
			DeskUser{Name: "user2"},
		}})
}
