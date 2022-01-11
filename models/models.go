package models

import (
	"database/sql"
	"time"
)

// Models is the wrapper for the database
type Models struct {
	DB DBModel
}

// NewModels returns Models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Instruction is the type for instructions
type Instruction struct {
	ID                  int            `json:"id"`
	Title               string         `json:"title"`
	Description         string         `json:"description"`
	Steps               []string       `json:"steps"`
	Rating              int            `json:"rating"`
	CreatedAt           time.Time      `json:"-"`
	UpdatedAt           time.Time      `json:"-"`
	InstructionCategory map[int]string `json:"categories"`
}

// Category is the type for categories
type Category struct {
	ID           int       `json:"-"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

// InstructionCategory is the type for instruction category
type InstructionCategory struct {
	ID            int       `json:"-"`
	InstructionID int       `json:"-"`
	CategoryID    int       `json:"-"`
	Category      Category  `json:"category"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}
