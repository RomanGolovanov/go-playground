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
