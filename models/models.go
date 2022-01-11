package models

import "time"

type Instruction struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Steps []string `json:"steps"`
	Rating int `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	InstructionCategory []Category `json:"-"`
}


type Category struct {
	ID int `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InstructionCategory struct {
	ID int `json:"id"`
	InstructionID int `json:"instruction_id"`
	CategoryID int `json:"category_id"`
	Category Category `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}