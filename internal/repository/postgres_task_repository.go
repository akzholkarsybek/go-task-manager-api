package repository

import (
	"database/sql"
	"github.com/Akakazkz/go-task-manager-api/internal/model"
)

type postgresTaskRepository struct{
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) TaskRepository{
	return &postgresTaskRepository{
		db: db, 
	}
}

func (r *postgresTaskRepository) Create(task *model.Task) error{
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
func (r *postgresTaskRepository) GetByID(id, userID int64) (*model.Task, error){
	panic("not implemented")
}

func (r *postgresTaskRepository) ListByUserID(userID int64) ([]*model.Task, error){
	rows, err := r.db.Query(`
		SELECT id, title, description, completed, due_date, user_id, created_at
		FROM tasks
		WHERE user_id = $1
		ORDER BY created_at DESC
	`,
	userID,
)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task

	for rows.Next(){
		var t model.Task
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Completed,
			&t.DueDate,
			&t.UserID,
			&t.CreatedAt,
		); err != nil{
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	if err := rows.Err(); err != nil {
    	return nil, err
	}
	return tasks, nil

}
func (r *postgresTaskRepository) Update(task *model.Task) error{
	panic("not implemented")
}
func (r *postgresTaskRepository) Delete(id, userID int64) error{
	panic("not implemented")
}


