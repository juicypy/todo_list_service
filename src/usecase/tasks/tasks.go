package tasks

import (
	"context"
	"github.com/juicypy/todo_list_service/src/entities"
	"time"
)

type TaskCommentsRepo interface {
	SearchByTask(ctx context.Context, userID string, taskID string) ([]entities.TaskComment, error)
}

type TasksRepo interface {
	UpsertTask(ctx context.Context, task entities.Task) error
	AllTasks(ctx context.Context, userID string) ([]entities.Task, error)
	TaskByID(ctx context.Context, userID string, taskID string) (entities.Task, error)
}

type LabelsRepo interface {
	LabelsByIDs(ctx context.Context, ids []string) ([]entities.Label, error)
}

type TasksUseCase struct {
	tasksRepo        TasksRepo
	labelsRepo       LabelsRepo
	taskCommentsRepo TaskCommentsRepo
}

func NewTasksUseCase(tasks TasksRepo, taskComments TaskCommentsRepo, labels LabelsRepo) *TasksUseCase {
	return &TasksUseCase{tasksRepo: tasks, taskCommentsRepo: taskComments, labelsRepo: labels}
}

func (s *TasksUseCase) CreateTask(ctx context.Context, task entities.Task) error {
	task.Status = entities.StatusTODO
	task.CreatedAt = time.Now().Unix()
	task.ModifiedAt = time.Now().Unix()

	return s.tasksRepo.UpsertTask(ctx, task)
}

func (s *TasksUseCase) AllTasks(ctx context.Context, userID string) ([]entities.TaskView, error) {
	tasks, err := s.tasksRepo.AllTasks(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]entities.TaskView, 0, len(tasks))
	for _, t := range tasks {
		comments, err := s.taskCommentsRepo.SearchByTask(ctx, userID, t.ID)
		if err != nil {
			return nil, err
		}

		labels, err := s.labelsRepo.LabelsByIDs(ctx, t.LabelIDs)
		if err != nil {
			return nil, err
		}

		res = append(res, entities.TaskView{
			Task:     t,
			Comments: comments,
			Labels:   labels,
		})
	}
	return res, nil
}
