// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: games.sql

package databases

import (
	"context"
	"time"
)

const createGame = `-- name: CreateGame :one
INSERT INTO games (
    id,
    hometeam_fullname,
    hometeam_shortname,
    hometeam_logourl,
    awayteam_fullname,
    awayteam_shortname,
    awayteam_logourl,
    channel,
    date,
    status,
    year,
    week,
    weekname,
    homeorunderline,
    homeorunderodd,
    awayoroverline,
    awayoroverodd,
    created_at_vegas,
    sportsbookid,
    homescore,
    awayscore,
    game_spread_winner,
    season_type,
    league
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15,
    $16,
    $17,
    $18,
    $19,
    $20,
    $21,
    $22,
    $23,
    $24
) RETURNING id, created_at, updated_at, hometeam_fullname, hometeam_shortname, hometeam_logourl, awayteam_fullname, awayteam_shortname, awayteam_logourl, channel, date, status, year, week, weekname, homeorunderline, homeorunderodd, awayoroverline, awayoroverodd, created_at_vegas, sportsbookid, homescore, awayscore, game_spread_winner, season_type, league
`

type CreateGameParams struct {
	ID                string       `json:"id"`
	HometeamFullname  string       `json:"hometeam_fullname"`
	HometeamShortname string       `json:"hometeam_shortname"`
	HometeamLogourl   string       `json:"hometeam_logourl"`
	AwayteamFullname  string       `json:"awayteam_fullname"`
	AwayteamShortname string       `json:"awayteam_shortname"`
	AwayteamLogourl   string       `json:"awayteam_logourl"`
	Channel           string       `json:"channel"`
	Date              time.Time    `json:"date"`
	Status            string       `json:"status"`
	Year              int32        `json:"year"`
	Week              int32        `json:"week"`
	Weekname          string       `json:"weekname"`
	Homeorunderline   float64      `json:"homeorunderline"`
	Homeorunderodd    int32        `json:"homeorunderodd"`
	Awayoroverline    float64      `json:"awayoroverline"`
	Awayoroverodd     int32        `json:"awayoroverodd"`
	CreatedAtVegas    string       `json:"created_at_vegas"`
	Sportsbookid      int32        `json:"sportsbookid"`
	Homescore         int32        `json:"homescore"`
	Awayscore         int32        `json:"awayscore"`
	GameSpreadWinner  SpreadWinner `json:"game_spread_winner"`
	SeasonType        SeasonType   `json:"season_type"`
	League            League       `json:"league"`
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (Game, error) {
	row := q.db.QueryRowContext(ctx, createGame,
		arg.ID,
		arg.HometeamFullname,
		arg.HometeamShortname,
		arg.HometeamLogourl,
		arg.AwayteamFullname,
		arg.AwayteamShortname,
		arg.AwayteamLogourl,
		arg.Channel,
		arg.Date,
		arg.Status,
		arg.Year,
		arg.Week,
		arg.Weekname,
		arg.Homeorunderline,
		arg.Homeorunderodd,
		arg.Awayoroverline,
		arg.Awayoroverodd,
		arg.CreatedAtVegas,
		arg.Sportsbookid,
		arg.Homescore,
		arg.Awayscore,
		arg.GameSpreadWinner,
		arg.SeasonType,
		arg.League,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HometeamFullname,
		&i.HometeamShortname,
		&i.HometeamLogourl,
		&i.AwayteamFullname,
		&i.AwayteamShortname,
		&i.AwayteamLogourl,
		&i.Channel,
		&i.Date,
		&i.Status,
		&i.Year,
		&i.Week,
		&i.Weekname,
		&i.Homeorunderline,
		&i.Homeorunderodd,
		&i.Awayoroverline,
		&i.Awayoroverodd,
		&i.CreatedAtVegas,
		&i.Sportsbookid,
		&i.Homescore,
		&i.Awayscore,
		&i.GameSpreadWinner,
		&i.SeasonType,
		&i.League,
	)
	return i, err
}

const deleteGame = `-- name: DeleteGame :exec
DELETE FROM games
WHERE id = $1
`

func (q *Queries) DeleteGame(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteGame, id)
	return err
}

const getGameByID = `-- name: GetGameByID :one
SELECT id, created_at, updated_at, hometeam_fullname, hometeam_shortname, hometeam_logourl, awayteam_fullname, awayteam_shortname, awayteam_logourl, channel, date, status, year, week, weekname, homeorunderline, homeorunderodd, awayoroverline, awayoroverodd, created_at_vegas, sportsbookid, homescore, awayscore, game_spread_winner, season_type, league FROM games
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetGameByID(ctx context.Context, id string) (Game, error) {
	row := q.db.QueryRowContext(ctx, getGameByID, id)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HometeamFullname,
		&i.HometeamShortname,
		&i.HometeamLogourl,
		&i.AwayteamFullname,
		&i.AwayteamShortname,
		&i.AwayteamLogourl,
		&i.Channel,
		&i.Date,
		&i.Status,
		&i.Year,
		&i.Week,
		&i.Weekname,
		&i.Homeorunderline,
		&i.Homeorunderodd,
		&i.Awayoroverline,
		&i.Awayoroverodd,
		&i.CreatedAtVegas,
		&i.Sportsbookid,
		&i.Homescore,
		&i.Awayscore,
		&i.GameSpreadWinner,
		&i.SeasonType,
		&i.League,
	)
	return i, err
}

const listGames = `-- name: ListGames :many
SELECT id, created_at, updated_at, hometeam_fullname, hometeam_shortname, hometeam_logourl, awayteam_fullname, awayteam_shortname, awayteam_logourl, channel, date, status, year, week, weekname, homeorunderline, homeorunderodd, awayoroverline, awayoroverodd, created_at_vegas, sportsbookid, homescore, awayscore, game_spread_winner, season_type, league FROM games
ORDER BY date
`

func (q *Queries) ListGames(ctx context.Context) ([]Game, error) {
	rows, err := q.db.QueryContext(ctx, listGames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Game
	for rows.Next() {
		var i Game
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.HometeamFullname,
			&i.HometeamShortname,
			&i.HometeamLogourl,
			&i.AwayteamFullname,
			&i.AwayteamShortname,
			&i.AwayteamLogourl,
			&i.Channel,
			&i.Date,
			&i.Status,
			&i.Year,
			&i.Week,
			&i.Weekname,
			&i.Homeorunderline,
			&i.Homeorunderodd,
			&i.Awayoroverline,
			&i.Awayoroverodd,
			&i.CreatedAtVegas,
			&i.Sportsbookid,
			&i.Homescore,
			&i.Awayscore,
			&i.GameSpreadWinner,
			&i.SeasonType,
			&i.League,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGamesByWeek = `-- name: ListGamesByWeek :many
SELECT id, created_at, updated_at, hometeam_fullname, hometeam_shortname, hometeam_logourl, awayteam_fullname, awayteam_shortname, awayteam_logourl, channel, date, status, year, week, weekname, homeorunderline, homeorunderodd, awayoroverline, awayoroverodd, created_at_vegas, sportsbookid, homescore, awayscore, game_spread_winner, season_type, league FROM games
WHERE week = $1 
AND year = $2 
AND season_type = $3
AND league = $4
`

type ListGamesByWeekParams struct {
	Week       int32      `json:"week"`
	Year       int32      `json:"year"`
	SeasonType SeasonType `json:"season_type"`
	League     League     `json:"league"`
}

func (q *Queries) ListGamesByWeek(ctx context.Context, arg ListGamesByWeekParams) ([]Game, error) {
	rows, err := q.db.QueryContext(ctx, listGamesByWeek,
		arg.Week,
		arg.Year,
		arg.SeasonType,
		arg.League,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Game
	for rows.Next() {
		var i Game
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.HometeamFullname,
			&i.HometeamShortname,
			&i.HometeamLogourl,
			&i.AwayteamFullname,
			&i.AwayteamShortname,
			&i.AwayteamLogourl,
			&i.Channel,
			&i.Date,
			&i.Status,
			&i.Year,
			&i.Week,
			&i.Weekname,
			&i.Homeorunderline,
			&i.Homeorunderodd,
			&i.Awayoroverline,
			&i.Awayoroverodd,
			&i.CreatedAtVegas,
			&i.Sportsbookid,
			&i.Homescore,
			&i.Awayscore,
			&i.GameSpreadWinner,
			&i.SeasonType,
			&i.League,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateGameAlreadyFound = `-- name: UpdateGameAlreadyFound :one
UPDATE games
    set hometeam_fullname = $2,
    hometeam_shortname = $3,
    hometeam_logourl = $4,
    awayteam_fullname = $5,
    awayteam_shortname = $6,
    awayteam_logourl = $7,
    channel = $8,
    date = $9,
    status = $10,
    year = $11,
    week = $12,
    weekname = $13,
    homeorunderline = $14,
    homeorunderodd = $15,
    awayoroverline = $16,
    awayoroverodd = $17,
    created_at_vegas = $18,
    sportsbookid = $19,
    homescore = $20,
    awayscore = $21,
    game_spread_winner = $22,
    season_type = $23,
    league = $24,
    updated_at = Now()
WHERE id = $1
RETURNING id, created_at, updated_at, hometeam_fullname, hometeam_shortname, hometeam_logourl, awayteam_fullname, awayteam_shortname, awayteam_logourl, channel, date, status, year, week, weekname, homeorunderline, homeorunderodd, awayoroverline, awayoroverodd, created_at_vegas, sportsbookid, homescore, awayscore, game_spread_winner, season_type, league
`

type UpdateGameAlreadyFoundParams struct {
	ID                string       `json:"id"`
	HometeamFullname  string       `json:"hometeam_fullname"`
	HometeamShortname string       `json:"hometeam_shortname"`
	HometeamLogourl   string       `json:"hometeam_logourl"`
	AwayteamFullname  string       `json:"awayteam_fullname"`
	AwayteamShortname string       `json:"awayteam_shortname"`
	AwayteamLogourl   string       `json:"awayteam_logourl"`
	Channel           string       `json:"channel"`
	Date              time.Time    `json:"date"`
	Status            string       `json:"status"`
	Year              int32        `json:"year"`
	Week              int32        `json:"week"`
	Weekname          string       `json:"weekname"`
	Homeorunderline   float64      `json:"homeorunderline"`
	Homeorunderodd    int32        `json:"homeorunderodd"`
	Awayoroverline    float64      `json:"awayoroverline"`
	Awayoroverodd     int32        `json:"awayoroverodd"`
	CreatedAtVegas    string       `json:"created_at_vegas"`
	Sportsbookid      int32        `json:"sportsbookid"`
	Homescore         int32        `json:"homescore"`
	Awayscore         int32        `json:"awayscore"`
	GameSpreadWinner  SpreadWinner `json:"game_spread_winner"`
	SeasonType        SeasonType   `json:"season_type"`
	League            League       `json:"league"`
}

func (q *Queries) UpdateGameAlreadyFound(ctx context.Context, arg UpdateGameAlreadyFoundParams) (Game, error) {
	row := q.db.QueryRowContext(ctx, updateGameAlreadyFound,
		arg.ID,
		arg.HometeamFullname,
		arg.HometeamShortname,
		arg.HometeamLogourl,
		arg.AwayteamFullname,
		arg.AwayteamShortname,
		arg.AwayteamLogourl,
		arg.Channel,
		arg.Date,
		arg.Status,
		arg.Year,
		arg.Week,
		arg.Weekname,
		arg.Homeorunderline,
		arg.Homeorunderodd,
		arg.Awayoroverline,
		arg.Awayoroverodd,
		arg.CreatedAtVegas,
		arg.Sportsbookid,
		arg.Homescore,
		arg.Awayscore,
		arg.GameSpreadWinner,
		arg.SeasonType,
		arg.League,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HometeamFullname,
		&i.HometeamShortname,
		&i.HometeamLogourl,
		&i.AwayteamFullname,
		&i.AwayteamShortname,
		&i.AwayteamLogourl,
		&i.Channel,
		&i.Date,
		&i.Status,
		&i.Year,
		&i.Week,
		&i.Weekname,
		&i.Homeorunderline,
		&i.Homeorunderodd,
		&i.Awayoroverline,
		&i.Awayoroverodd,
		&i.CreatedAtVegas,
		&i.Sportsbookid,
		&i.Homescore,
		&i.Awayscore,
		&i.GameSpreadWinner,
		&i.SeasonType,
		&i.League,
	)
	return i, err
}

const updateGameOdds = `-- name: UpdateGameOdds :one
UPDATE games 
    set homeorunderline = $2,
    homeorunderodd = $3,
    awayoroverline = $4,
    awayoroverodd = $5,
    game_spread_winner = $6,
    updated_at = Now()
WHERE id = $1
RETURNING id, created_at, updated_at, hometeam_fullname, hometeam_shortname, hometeam_logourl, awayteam_fullname, awayteam_shortname, awayteam_logourl, channel, date, status, year, week, weekname, homeorunderline, homeorunderodd, awayoroverline, awayoroverodd, created_at_vegas, sportsbookid, homescore, awayscore, game_spread_winner, season_type, league
`

type UpdateGameOddsParams struct {
	ID               string       `json:"id"`
	Homeorunderline  float64      `json:"homeorunderline"`
	Homeorunderodd   int32        `json:"homeorunderodd"`
	Awayoroverline   float64      `json:"awayoroverline"`
	Awayoroverodd    int32        `json:"awayoroverodd"`
	GameSpreadWinner SpreadWinner `json:"game_spread_winner"`
}

func (q *Queries) UpdateGameOdds(ctx context.Context, arg UpdateGameOddsParams) (Game, error) {
	row := q.db.QueryRowContext(ctx, updateGameOdds,
		arg.ID,
		arg.Homeorunderline,
		arg.Homeorunderodd,
		arg.Awayoroverline,
		arg.Awayoroverodd,
		arg.GameSpreadWinner,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HometeamFullname,
		&i.HometeamShortname,
		&i.HometeamLogourl,
		&i.AwayteamFullname,
		&i.AwayteamShortname,
		&i.AwayteamLogourl,
		&i.Channel,
		&i.Date,
		&i.Status,
		&i.Year,
		&i.Week,
		&i.Weekname,
		&i.Homeorunderline,
		&i.Homeorunderodd,
		&i.Awayoroverline,
		&i.Awayoroverodd,
		&i.CreatedAtVegas,
		&i.Sportsbookid,
		&i.Homescore,
		&i.Awayscore,
		&i.GameSpreadWinner,
		&i.SeasonType,
		&i.League,
	)
	return i, err
}
