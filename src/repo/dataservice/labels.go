package dataservice

import (
	"context"
	"github.com/doug-martin/goqu/v9"
	"github.com/juicypy/todo_list_service/src/entities"
)

const labelsTableName = "label"

type LabelsRepo struct {
	db *goqu.Database
}

func NewLabelsRepo(db *goqu.Database) *LabelsRepo {
	return &LabelsRepo{
		db: db,
	}
}

func (s *LabelsRepo) LabelsByIDs(ctx context.Context, ids []string) ([]entities.Label, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	result := make([]entities.Label, 0)
	err := s.db.From(labelsTableName).Where(goqu.I("id").In(ids)).Executor().ScanStructsContext(ctx, &result)
	return result, err
}
