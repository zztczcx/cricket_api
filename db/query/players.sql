-- name: GetPlayerOfMostRuns :one
SELECT * FROM players
ORDER BY runs DESC LIMIT 1;

-- name: GetPlayerOfMostRunsByCareerEndYear :one
SELECT * FROM players
WHERE career_end_year = ?
ORDER BY runs DESC LIMIT 1;

-- name: GetPlayersByCareerYear :many
SELECT * FROM players
WHERE career_start_year <= sqlc.arg(career_year) AND career_end_year >= sqlc.arg(career_year)
ORDER By id ASC;

-- name: CreatePlayer :execresult
INSERT INTO players (
    name,
    career_start_year,
    career_end_year,
    matches,
    inns,
    not_outs,
    runs,
    highest_scores,
    average,
    faced_balls,
    strike_rate,
    score_hundreds,
    score_fiftys,
    score_zeros
) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: DeleteAllPlayers :execresult
DELETE FROM players;
