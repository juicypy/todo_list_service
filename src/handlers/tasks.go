package handlers

import (
	"context"
	"encoding/json"
	"github.com/juicypy/todo_list_service/src/entities"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type TasksUseCase interface {
	CreateTask(ctx context.Context, task entities.Task) error
	AllTasks(ctx context.Context, userID string) ([]entities.TaskView, error)
}

func NewTasksHandler(uc TasksUseCase, logger zap.SugaredLogger) *TasksHandler {
	return &TasksHandler{
		uc:     uc,
		logger: logger,
	}
}

type TasksHandler struct {
	uc     TasksUseCase
	logger zap.SugaredLogger
}

func (s *TasksHandler) Create(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := s.logger.With("where", "TasksHandler.Create")

	user, err := userFromCtx(ctx)
	if err != nil {
		log.Errorw("error reading user from context", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warnw("error reading body", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}

	task := &entities.Task{}
	err = json.Unmarshal(data, task)
	if err != nil {
		log.Warnw("error unmarshalling body", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}

	task.UserID = user.ID
	err = s.uc.CreateTask(ctx, *task)
	if err != nil {
		log.Errorw("error creating a task", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (s *TasksHandler) GetAll(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := s.logger.With("where", "TasksHandler.GetAll")

	user, err := userFromCtx(ctx)
	if err != nil {
		log.Errorw("error reading user from context", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}

	res, err := s.uc.AllTasks(ctx, user.ID)
	if err != nil {
		log.Errorw("error retrieving all tasks", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}

	renderJSON(rw, http.StatusOK, res)
}
