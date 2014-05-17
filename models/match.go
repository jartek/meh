package models

type Match struct {
	Id         int
	HomeTeamId int `db:"home_team_id"`
	AwayTeamId int `db:"away_team_id"`
	StadiumId  int `db:"stadium_id"`
}
