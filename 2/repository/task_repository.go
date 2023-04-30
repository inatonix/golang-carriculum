package repository

import (
	"2/model"
	"database/sql"
	"fmt"
)

type TaskRepository interface {
	Create(task *model.Task) (int, error)
	Read(id int) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
}

type taskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *taskRepositoryImpl {
	return &taskRepositoryImpl{db: db}
}

func (r *taskRepositoryImpl) Create(task *model.Task) (int, error) {
	stmt := `INSERT INTO tasks (title) VALUES (?)`
	result, err := r.db.Exec(stmt, task.Title)
	if err != nil {
		return 0, err
	}
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println(lastInsertedID)
	return int(lastInsertedID), err
}

func (r *taskRepositoryImpl) Read(id int) (*model.Task, error) {
	stmt := `SELECT id, title FROM tasks WHERE id = ?`
	task := model.Task{}
	err := r.db.QueryRow(stmt, id).Scan(&task.ID, &task.Title)
	return &task, err
}

func (r *taskRepositoryImpl) Update(task *model.Task) error {
	stmt := `UPDATE tasks SET title = ? WHERE id = ?`
	rows, err := r.db.Exec(stmt, task.Title, task.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (r *taskRepositoryImpl) Delete(id int) error {
	stmt := `DELETE FROM tasks WHERE id = ?`
	rows, err := r.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
