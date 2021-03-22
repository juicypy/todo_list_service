package entities

type TaskComment struct {
	ID        string `json:"id" db:"id" goqu:"skipinsert,skipupdate"`
	UserID    string `json:"-" db:"user_id"`
	TaskID    string `json:"task_id" db:"task_id"`
	Comment   string `json:"comment" db:"comment"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
}
