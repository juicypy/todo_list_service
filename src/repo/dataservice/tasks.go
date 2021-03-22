package dataservice

import (
	"context"

	"github.com/doug-martin/goqu/v9"

	"github.com/juicypy/todo_list_service/src/entities"
)

const taskTableName = "task"

type TaskDBRepo struct {
	db *goqu.Database
}

func NewTaskDBRepo(db *goqu.Database) *TaskDBRepo {
	return &TaskDBRepo{
		db: db,
	}
}

func (s *TaskDBRepo) UpsertTask(ctx context.Context, task entities.Task) error {
	_, err := s.db.Insert(taskTableName).Rows(task).
		OnConflict(goqu.DoUpdate("id", task)).
		Executor().
		ExecContext(ctx)
	return err
}

func (s *TaskDBRepo) AllTasks(ctx context.Context, userID string) ([]entities.Task, error) {
	result := make([]entities.Task, 0)
	err := s.db.From(taskTableName).Where(goqu.Ex{"user_id": userID}).Executor().ScanStructsContext(ctx, &result)
	return result, err
}

func (s *TaskDBRepo) TaskByID(ctx context.Context, userID string, taskID string) (entities.Task, error) {
	result := entities.Task{}
	err := s.db.From(taskTableName).Where(goqu.Ex{"id": taskID, "user_id": userID}).Executor().ScanStructsContext(ctx, &result)
	return result, err
}
