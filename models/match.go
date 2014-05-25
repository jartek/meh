package models

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
