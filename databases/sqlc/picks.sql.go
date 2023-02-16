// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: picks.sql

package databases

import (
	"context"
)

const createPick = `-- name: CreatePick :one
INSERT INTO picks (
    username,
    game_id,
    year, --pickems league year
    week, --pickems league week
    league, --idk if we need
    user_pick, -- need this
    user_pick_type, -- favorite or underdog
    game_spread_winner -- from game_id
    ) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
    ) RETURNING username, game_id, created_at, updated_at, year, week, league, user_pick, user_pick_type, game_spread_winner
`

type CreatePickParams struct {
	Username         string       `json:"username"`
	GameID           string       `json:"game_id"`
	Year             int32        `json:"year"`
	Week             int32        `json:"week"`
	League           League       `json:"league"`
	UserPick         SpreadPicks  `json:"user_pick"`
	UserPickType     SpreadType   `json:"user_pick_type"`
	GameSpreadWinner SpreadWinner `json:"game_spread_winner"`
}

func (q *Queries) CreatePick(ctx context.Context, arg CreatePickParams) (Pick, error) {
	row := q.db.QueryRowContext(ctx, createPick,
		arg.Username,
		arg.GameID,
		arg.Year,
		arg.Week,
		arg.League,
		arg.UserPick,
		arg.UserPickType,
		arg.GameSpreadWinner,
	)
	var i Pick
	err := row.Scan(
		&i.Username,
		&i.GameID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Year,
		&i.Week,
		&i.League,
		&i.UserPick,
		&i.UserPickType,
		&i.GameSpreadWinner,
	)
	return i, err
}
