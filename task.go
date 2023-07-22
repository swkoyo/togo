package main

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	DueAt       *time.Time `json:"due_at"`
	IsComplete  bool       `json:"is_complete"`
	CategoryID  *int       `json:"category_id"`
	Category    *Category  `json:"category"`
}

func scanTask(rows *sql.Rows) (*Task, error) {
	var t Task
	var tDescription sql.NullString
	var tDueAt sql.NullString
	var cID sql.NullInt64
	var cName sql.NullString

	if err := rows.Scan(
		&t.ID,
		&t.Name,
		&tDescription,
		&tDueAt,
		&t.IsComplete,
		&cID,
		&cName,
	); err != nil {
		return nil, err
	}

	if tDescription.Valid {
		t.Description = &tDescription.String
	}

	if tDueAt.Valid {
		parsedDate, err := time.Parse("2006-01-02 15:04:05", tDueAt.String)
		if err != nil {
			return nil, err
		}
		t.DueAt = &parsedDate
	}

	if cID.Valid && cName.Valid {
		convertedID := int(cID.Int64)
		t.CategoryID = &convertedID
		t.Category = &Category{
			ID:   convertedID,
			Name: cName.String,
		}
	}

	return &t, nil
}

func GetTask(id int) (*Task, error) {
	rows, err := DB.Query(`
        SELECT
            t.id as id,
            t.name as name,
            t.description as description,
            t.due_at as due_at,
            t.is_complete as is_complete,
            t.category_id as category_id,
            c.id as cat_id,
            c.name as category_name,
        FROM tasks as t
        LEFT JOIN categories as c
            ON c.id = t.category_id
        WHERE t.id = ?
    `, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var t *Task
	if rows.Next() {
		t, err = scanTask(rows)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func GetTasks() ([]Task, error) {
	rows, err := DB.Query(`
        SELECT
            t.id as id,
            t.name as name,
            t.description as description,
            t.due_at as due_at,
            t.is_complete as is_complete,
            t.category_id as category_id,
            c.name as category_name
        FROM tasks as t
        LEFT JOIN categories as c
            ON c.id = t.category_id
        ORDER BY t.id DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		t, err := scanTask(rows)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, *t)
	}
	return tasks, nil
}

func (t *Task) Create() error {
	_, err := DB.Exec(`
        INSERT INTO tasks (
            name,
            description,
            due_at,
            is_complete,
            category_id
        ) VALUES (?, ?, ?, ?, ?, ?)
    `, t.Name, t.Description, t.DueAt, t.IsComplete, t.CategoryID)
	return err
}

func (t *Task) Update() error {
	_, err := DB.Exec(`
        UPDATE tasks SET
            name = ?,
            description = ?,
            created_at = ?,
            due_at = ?,
            is_complete = ?,
            category_id = ?
        WHERE id = ?
    `, t.Name, t.Description, t.DueAt, t.IsComplete, t.CategoryID, t.ID)
	return err
}

func (t *Task) Delete() error {
	_, err := DB.Exec(`DELETE FROM tasks WHERE id = ?`, t.ID)
	return err
}
