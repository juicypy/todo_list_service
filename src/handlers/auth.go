package handlers

import (
	"context"
	"github.com/juicypy/todo_list_service/src/entities"
	"go.uber.org/zap"
	"net/http"
)

type AuthUsecase interface {
	UserByID(ctx context.Context, guid string) (*entities.UserDB, error)
}

func NewAuthHandler(uc AuthUsecase, logger zap.SugaredLogger) *AuthHandler {
	return &AuthHandler{
		uc:     uc,
		logger: logger,
	}
}

type AuthHandler struct {
	uc     AuthUsecase
	logger zap.SugaredLogger
}

func (s *AuthHandler) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := s.logger.With("where", "AuthHandler.CheckAuth")

		userID, err := token(*r)
		if err != nil {
			log.Warnw("invalid auth header", "err", err, "header", r.Header.Get(authHeader))
			sendBadRequest(rw, "error_400", "invalid header")
			return
		}

		user, err := s.uc.UserByID(ctx, userID)
		if err != nil {
			log.Warnw("error retrieving user db info", "err", err)
			sendForbidden(rw, "error_403", "invalid token")
			return
		}

		r = r.WithContext(context.WithValue(ctx, userCtxKey, user))

		next.ServeHTTP(rw, r)
	})
}
