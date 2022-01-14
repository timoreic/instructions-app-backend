package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
)

type DBModel struct {
	DB *sql.DB
}

// Get returns one instruction and error, if any
func (m *DBModel) Get(id int) (*Instruction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, steps, rating, created_at, updated_at from instructions where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var instruction Instruction

	err := row.Scan(
		&instruction.ID,
		&instruction.Title,
		&instruction.Description,
		pq.Array(&instruction.Steps),
		&instruction.Rating,
		&instruction.CreatedAt,
		&instruction.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Get categories, if any
	query = `select
					    ic.id, ic.instruction_id, ic.category_id, c.category_name
					from 
							instructions_categories ic
							left join categories c on (c.id = ic.category_id)
					where
							ic.instruction_id = $1
	`
	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	categories := make(map[int]string)
	for rows.Next() {
		var ic InstructionCategory
		err := rows.Scan(
			&ic.ID,
			&ic.InstructionID,
			&ic.CategoryID,
			&ic.Category.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		categories[ic.ID] = ic.Category.CategoryName
	}

	instruction.InstructionCategory = categories

	return &instruction, nil
}

// All returns all instructions and error, if any
func (m *DBModel) All(category ...int) ([]*Instruction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""
	if len(category) > 0 {
		where = fmt.Sprintf("where id in (select instruction_id from instructions_categories where category_id = %d)", category[0])
	}

	query := fmt.Sprintf(`select id, title, description, steps, rating, created_at, updated_at from instructions  %s order by title`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var instructions []*Instruction

	for rows.Next() {
		var instruction Instruction
		err := rows.Scan(
			&instruction.ID,
			&instruction.Title,
			&instruction.Description,
			pq.Array(&instruction.Steps),
			&instruction.Rating,
			&instruction.CreatedAt,
			&instruction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Get categories, if any
		categoryQuery := `select
					    ic.id, ic.instruction_id, ic.category_id, c.category_name
					from 
							instructions_categories ic
							left join categories c on (c.id = ic.category_id)
					where
							ic.instruction_id = $1
	`
		categoryRows, _ := m.DB.QueryContext(ctx, categoryQuery, instruction.ID)

		categories := make(map[int]string)
		for categoryRows.Next() {
			var ic InstructionCategory
			err := categoryRows.Scan(
				&ic.ID,
				&ic.InstructionID,
				&ic.CategoryID,
				&ic.Category.CategoryName,
			)
			if err != nil {
				return nil, err
			}
			categories[ic.ID] = ic.Category.CategoryName
		}
		categoryRows.Close()

		instruction.InstructionCategory = categories
		instructions = append(instructions, &instruction)
	}
	return instructions, nil
}

// CategoriesAll returns all categories and error, if any
func (m *DBModel) CategoriesAll() ([]*Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, category_name, created_at, updated_at from categories order by category_name`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*Category

	for rows.Next() {
		var c Category
		err := rows.Scan(
			&c.ID,
			&c.CategoryName,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &c)
	}

	return categories, nil
}
