package dataservice

import (
	"context"
	"github.com/doug-martin/goqu/v9"
	"github.com/juicypy/todo_list_service/src/entities"
)

const userTableName = "user"

type UserDBRepo struct {
	db *goqu.Database
}

func NewUserDBRepo(db *goqu.Database) *UserDBRepo {
	return &UserDBRepo{
		db: db,
	}
}

func (s *UserDBRepo) UpsertUser(ctx context.Context, user entities.UserDB) (string, error) {
	var id string
	_, err := s.db.Insert(userTableName).Returning(goqu.C("id")).
		Rows(user).OnConflict(goqu.DoUpdate("id", user)).
		Executor().ScanValContext(ctx, &id)
	return id, err
}

func (s *UserDBRepo) UserByID(ctx context.Context, id string) (bool, *entities.UserDB, error) {
	user := &entities.UserDB{}
	found, err := s.db.From(userTableName).Where(goqu.Ex{"id": id}).ScanStructContext(ctx, user)
	return found, user, err
}
