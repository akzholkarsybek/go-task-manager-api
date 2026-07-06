package repository

import (
	"database/sql"
	"errors"

	"github.com/Akakazkz/go-task-manager-api/internal/model"
)

type postgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) TaskRepository {
	return &postgresTaskRepository{
		db: db,
	}
}

func (r *postgresTaskRepository) Create(task *model.Task) error {
	query := `
		INSERT INTO tasks ( title, description, completed, due_date, user_id, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id
	`
	return r.db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.DueDate,
		task.UserID,
		task.CreatedAt,
	).Scan(&task.ID)
}
func (r *postgresTaskRepository) GetByID(id, userID int64) (*model.Task, error) {
	var task model.Task
	query := `
		SELECT id, title, description, completed, due_date, user_id, created_at 
		FROM tasks
		WHERE id = $1
		AND user_id = $2
	`
	err := r.db.QueryRow(query, id, userID).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.DueDate,
		&task.UserID,
		&task.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *postgresTaskRepository) ListByUserID(userID int64) ([]*model.Task, error) {
	rows, err := r.db.Query(`
		SELECT id, title, description, completed, due_date, user_id, created_at
		FROM tasks
		WHERE user_id = $1
		ORDER BY created_at DESC
	`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task

	for rows.Next() {
		var t model.Task
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Completed,
			&t.DueDate,
			&t.UserID,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil

}
func (r *postgresTaskRepository) Update(task *model.Task) error {
	result, err := r.db.Exec(`
	UPDATE tasks
	SET
		title = $1,
		description = $2,
		completed = $3,
		due_date = $4
	WHERE id = $5
	AND user_id = $6
	`,
		task.Title,
		task.Description,
		task.Completed,
		task.DueDate,
		task.ID,
		task.UserID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("task not found")
	}
	return nil

}
func (r *postgresTaskRepository) Delete(id, userID int64) error {
	results, err := r.db.Exec(`
	DELETE FROM tasks
	WHERE id = $1
	AND user_id = $2
	`,
	id,
	userID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0{
		return errors.New("task not found")
	}
	return nil
}
