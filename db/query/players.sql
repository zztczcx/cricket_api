-- name: GetPlayersOfMostRuns :one
SELECT * FROM players
ORDER BY runs DESC LIMIT 1;

-- name: GetPlayersOfMostRunsByCareerEndYear :one
SELECT * FROM players
WHERE career_end_year = ?
ORDER BY runs DESC LIMIT 1;

-- name: GetPlayersByCareerYear :many
SELECT * FROM players
WHERE career_start_year <= sqlc.arg(career_year) AND career_end_year >= sqlc.arg(career_year);

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
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
);

