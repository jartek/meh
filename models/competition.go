package models

type Competition struct {
	Id        int64
	Name      string
	CreatedAt int64 `db:"created_at"`
	UpdatedAt int64 `db:"updated_at"`
}
