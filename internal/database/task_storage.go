package database

import (
	"HomeWork1/internal/entity"
	"database/sql"
	"errors"
	"fmt"
)

type TaskStorage struct {
	DB *sql.DB
}

func NewTaskStorage(db *sql.DB) *TaskStorage {
	return &TaskStorage{DB: db}
}

func (ts *TaskStorage) Get(id string) (*entity.Task, error) {
	query := "SELECT id, status, result FROM tasks WHERE id = $1"

	row := ts.DB.QueryRow(query, id)

	var task entity.Task

	err := row.Scan(&task.ID, &task.Status, &task.Result)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("task with id(%s) is not found", id)
		}
		return nil, fmt.Errorf("unable to scan task: %w", err)
	}

	return &task, err
}

func (ts *TaskStorage) Put(task entity.Task) error {
	query := "UPDATE tasks SET status = $2, result = $3 WHERE id = $1"

	result, err := ts.DB.Exec(query, task.ID, task.Status, task.Result)

	if err != nil {
		return fmt.Errorf("unable to update task with id(%s): %w", task.ID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")

	}
	return nil
}

func (ts *TaskStorage) Post(task entity.Task) error {
	query := "INSERT INTO tasks (id, status, result) VALUES($1, $2, $3)"

	result, err := ts.DB.Exec(query, task.ID, task.Status, task.Result)
	if err != nil {
		return fmt.Errorf("unable to insert new user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were inserted")
	}

	return nil
}

func (ts *TaskStorage) Delete(id string) error {
	query := "DELETE FROM tasks WHERE id = $1"
	//TODO: INSERT DRY CONCEPT
	result, err := ts.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("unable to insert new user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were inserted")
	}

	return nil
}
