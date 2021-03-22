package handlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/juicypy/todo_list_service/src/entities"
)

type TaskCommentsRepo interface {
	Create(ctx context.Context, task entities.TaskComment) error
	Delete(ctx context.Context, userID, taskID, commentID string) error
}

func NewTaskCommentsHandler(repo TaskCommentsRepo, logger zap.SugaredLogger) *TasksCommentsHandler {
	return &TasksCommentsHandler{
		repo:   repo,
		logger: logger,
	}
}

type TasksCommentsHandler struct {
	repo   TaskCommentsRepo
	logger zap.SugaredLogger
}

func (s *TasksCommentsHandler) Create(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := s.logger.With("where", "TasksCommentsHandler.Create")

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

	comment := &entities.TaskComment{}
	err = json.Unmarshal(data, comment)
	if err != nil {
		log.Warnw("error unmarshalling body", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}

	comment.UserID = user.ID
	comment.CreatedAt = time.Now().Unix()
	err = s.repo.Create(ctx, *comment)
	if err != nil {
		log.Errorw("error creating a comment", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (s *TasksCommentsHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := s.logger.With("where", "TasksCommentsHandler.Delete")

	user, err := userFromCtx(ctx)
	if err != nil {
		log.Errorw("error reading user from context", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}

	params := mux.Vars(r)
	taskID := params["task_id"]
	commentID := params["comment_id"]

	err = s.repo.Delete(ctx, user.ID, taskID, commentID)
	if err != nil {
		log.Errorw("error deleting a task attachment", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
