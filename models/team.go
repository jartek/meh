package models

type Team struct {
	Id       int
	Name     string
	NickName string `db:"nick_name"`
}
