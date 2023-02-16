-- name: CreateGame :one
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
) RETURNING *;

-- name: GetGameByID :one
SELECT * FROM games
WHERE id = $1 LIMIT 1;

-- name: ListGamesByWeek :many
SELECT * FROM games
WHERE week = $1 
AND year = $2 
AND season_type = $3
AND league = $4;

-- name: ListGames :many
SELECT * FROM games
ORDER BY date;

-- name: UpdateGameOdds :one
UPDATE games 
    set homeorunderline = $2,
    homeorunderodd = $3,
    awayoroverline = $4,
    awayoroverodd = $5,
    game_spread_winner = $6,
    updated_at = Now()
WHERE id = $1
RETURNING *;

-- name: UpdateGameAlreadyFound :one
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
RETURNING *;

-- name: DeleteGame :exec
DELETE FROM games
WHERE id = $1;