-- name: CreatePick :one
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
    ) RETURNING *;