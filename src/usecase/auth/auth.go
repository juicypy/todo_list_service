package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	"github.com/juicypy/todo_list_service/src/entities"
)

const (
	saltByteSize = 24
)

type UserDBRepo interface {
	UpsertUser(ctx context.Context, user entities.UserDB) (string, error)
	UserByID(ctx context.Context, id string) (bool, *entities.UserDB, error)
}

type AuthUsecase struct {
	repo UserDBRepo
}

func NewAuthUsecase(db UserDBRepo) *AuthUsecase {
	return &AuthUsecase{
		repo: db,
	}
}

func (s *AuthUsecase) SignUp(ctx context.Context, user *entities.UserToCreate) (id string, err error) {
	salt := make([]byte, saltByteSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}

	saltString := base64.StdEncoding.EncodeToString(salt)
	salted := fmt.Sprintf("%s%s", user.Password, saltString)
	hash := sha256.Sum256([]byte(salted))

	userDB := user.ToUserDB()
	userDB.PasswordDB = base64.StdEncoding.EncodeToString(hash[:])
	userDB.Salt = saltString

	return s.repo.UpsertUser(ctx, userDB)
}

func (s *AuthUsecase) UserByID(ctx context.Context, id string) (*entities.UserDB, error) {
	found, user, err := s.repo.UserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if !found || user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
