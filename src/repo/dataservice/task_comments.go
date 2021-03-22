package dataservice

import (
	"context"

	"github.com/doug-martin/goqu/v9"

	"github.com/juicypy/todo_list_service/src/entities"
)

const taskCommentsTableName = "task_comments"

type TaskCommentsRepo struct {
	db *goqu.Database
}

func NewTaskCommentsRepo(db *goqu.Database) *TaskCommentsRepo {
	return &TaskCommentsRepo{
		db: db,
	}
}

func (s *TaskCommentsRepo) Create(ctx context.Context, comment entities.TaskComment) error {
	_, err := s.db.Insert(taskCommentsTableName).Rows(comment).
		Executor().ExecContext(ctx)
	return err
}

func (s *TaskCommentsRepo) SearchByTask(ctx context.Context, userID string, taskID string) ([]entities.TaskComment, error) {
	result := make([]entities.TaskComment, 0)
	err := s.db.From(taskCommentsTableName).Where(goqu.Ex{"user_id": userID, "task_id": taskID}).Executor().ScanStructsContext(ctx, &result)
	return result, err
}

func (s *TaskCommentsRepo) Delete(ctx context.Context, userID, taskID, commentID string) error {
	_, err := s.db.Delete(taskCommentsTableName).
		Where(goqu.Ex{"id": commentID, "user_id": userID, "task_id": taskID}).
		Executor().ExecContext(ctx)

	return err
}
