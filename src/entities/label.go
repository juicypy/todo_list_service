package entities

type Label struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Color string `json:"color" db:"color"`
}
