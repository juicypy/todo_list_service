package entities

import "github.com/lib/pq"

type TaskStatus int

const (
	StatusTODO TaskStatus = iota
	StatusInProgress
	StatusDone
	StatusArchived
)

type Task struct {
	ID          string         `json:"id" db:"id" goqu:"skipinsert,skipupdate"`
	UserID      string         `json:"-" db:"user_id"`
	Name        string         `json:"name" db:"name"`
	Status      TaskStatus     `json:"status" db:"status"`
	Description string         `json:"description" db:"description"`
	DateFrom    int64          `json:"date_from" db:"date_from"`
	DateTo      int64          `json:"date_to" db:"date_to"`
	LabelIDs    pq.StringArray `json:"label_ids" db:"label_ids"`
	ModifiedAt  int64          `json:"modified_at" db:"modified_at"`
	CreatedAt   int64          `json:"created_at" db:"created_at" goqu:"skipupdate"`
}

type TaskView struct {
	Task
	Labels   []Label       `json:"labels"`
	Comments []TaskComment `json:"comments"`
}
