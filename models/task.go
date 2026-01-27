package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID			uuid.UUID	`json:"id"`
	Title		string		`json:"title"`
	Done		bool		`json:"done"`
	CreatedAt	time.Time	`json:"created_at"`		
	// Desc		string		`json:"description,omitempty"`
}
