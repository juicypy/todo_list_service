package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"

	"github.com/juicypy/todo_list_service/src/entities"
)

const (
	authHeader = "Authorization"

	headerErrorMsg = "empty_authorization_header"
	userCtxKey     = iota
)

type UserUseCase interface {
	SignUp(ctx context.Context, user *entities.UserToCreate) (id string, err error)
}

func NewUserHandler(uc UserUseCase, logger zap.SugaredLogger) *UserHandler {
	return &UserHandler{
		uc:     uc,
		logger: logger,
	}
}

type UserHandler struct {
	uc     UserUseCase
	logger zap.SugaredLogger
}

func (s *UserHandler) SignUp(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := s.logger.With("where", "UserHandler.SignUp")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warnw("error reading body", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}

	usr := &entities.UserToCreate{}
	err = json.Unmarshal(data, usr)
	if err != nil {
		log.Warnw("error unmarshalling body", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}

	id, err := s.uc.SignUp(ctx, usr)
	if err != nil {
		log.Errorw("error creating a user", "err", err)
		sendInternalServerError(rw, "error_500", "internal error")
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(id))
}
