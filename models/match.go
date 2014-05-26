package models

import (
	"database/sql"
	"strconv"

	"github.com/go-martini/martini"
)

// TYPES
// 1 - Group
// 2 - Knockout
// 3 - League

type Match struct {
	Id            int64
	HomeTeamId    int64 `db:"home_team_id"`
	AwayTeamId    int64 `db:"away_team_id"`
	StadiumId     int64 `db:"stadium_id"`
	Type          int64 // Will hold whether the match is a league/knockout/etc
	HomeTeamScore int64 `db:"home_team_score"`
	AwayTeamScore int64 `db:"away_team_score"`
	MatchTime     int64 `db:"match_time"`
	CreatedAt     int64 `db:"created_at"`
	UpdatedAt     int64 `db:"updated_at"`
	CompetitionId int64 `db:"competition_id"`
}

func GetAllMatches(db *sql.DB) ([]interface{}, error) {
	return GetAll(db, &Match{})
}

func GetMatch(db *sql.DB, params martini.Params) (interface{}, error) {
	id, _ := strconv.Atoi(params["id"])
	return GetOne(db, &Match{}, id)
}
