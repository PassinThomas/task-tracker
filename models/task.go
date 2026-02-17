package models

import (
	"time"
)

type Task struct {
	ID			int			`json:"id"`
	Title		string		`json:"title"`
	Done		bool		`json:"done"`
	CreatedAt	time.Time	`json:"created_at"`		
	UpdatedAt	time.Time	`json:"updated_at"`
}

type FlgUpdate struct {
	Done		bool
	NotDone		bool
	NewTitle	string
}

type FilterOptions struct {
	Done		bool
	Undone		bool
	Updated		bool
	NotUpdated	bool
}

type ListOptions struct {
	Filter  FilterOptions
	Sort    string
	Order   string
}
