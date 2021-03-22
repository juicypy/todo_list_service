package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/juicypy/todo_list_service/src/entities"
)

// Message error message
type Message struct {
	ErrName        string `json:"error"`
	ErrDescription string `json:"description"`
}

func token(r http.Request) (string, error) {
	bearer := r.Header.Get(authHeader)
	if bearer == "" {
		return "", errors.New(headerErrorMsg)
	}
	tokenRaw := strings.Split(bearer, " ")
	if len(tokenRaw) != 2 {
		return "", errors.New(headerErrorMsg)
	}
	return tokenRaw[1], nil
}

func userFromCtx(ctx context.Context) (*entities.UserDB, error) {
	usrI := ctx.Value(userCtxKey)
	usr, ok := usrI.(*entities.UserDB)
	if !ok {
		return nil, errors.New("could not get user data from context")
	}
	return usr, nil
}

func sendStatusCodeWithMessage(w http.ResponseWriter, code int, errName, errDescription string) {
	data, err := json.Marshal(Message{ErrName: errName, ErrDescription: errDescription})
	if err != nil {
		code = http.StatusInternalServerError
		data = []byte(`{"error": "serializing failed"}`)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(data)
}

// RenderJSON render json content type with status code
func renderJSON(w http.ResponseWriter, status int, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		status = http.StatusInternalServerError
		data = []byte(`{"error": "serializing failed"}`)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(data)
}

func sendBadRequest(w http.ResponseWriter, errName, errDescription string) {
	sendStatusCodeWithMessage(w, http.StatusBadRequest, errName, errDescription)
}

func sendInternalServerError(w http.ResponseWriter, errName, errDescription string) {
	sendStatusCodeWithMessage(w, http.StatusInternalServerError, errName, errDescription)
}

func sendForbidden(w http.ResponseWriter, errName, errDescription string) {
	sendStatusCodeWithMessage(w, http.StatusForbidden, errName, errDescription)
}
