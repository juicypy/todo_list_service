package entities

type UserRole int

type UserToCreate struct {
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password"`
}

func (s UserToCreate) ToUserDB() UserDB {
	return UserDB{
		Name:  s.Name,
		Email: s.Email,
		Phone: s.Phone,
	}
}

type UserDB struct {
	ID         string `json:"-" db:"id" goqu:"skipinsert,skipupdate"`
	Name       string `json:"name" db:"name"`
	Email      string `json:"email" db:"email"`
	Phone      string `json:"phone" db:"phone"`
	PasswordDB string `json:"-" db:"password"`
	Salt       string `json:"-" db:"salt"`
}
