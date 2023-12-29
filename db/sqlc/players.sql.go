// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: players.sql

package db

import (
	"context"
	"database/sql"
)

const createPlayer = `-- name: CreatePlayer :execresult
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
) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreatePlayerParams struct {
	Name            string          `json:"name"`
	CareerStartYear sql.NullInt64   `json:"career_start_year"`
	CareerEndYear   sql.NullInt64   `json:"career_end_year"`
	Matches         sql.NullInt64   `json:"matches"`
	Inns            sql.NullInt64   `json:"inns"`
	NotOuts         sql.NullInt64   `json:"not_outs"`
	Runs            sql.NullInt64   `json:"runs"`
	HighestScores   sql.NullInt64   `json:"highest_scores"`
	Average         sql.NullFloat64 `json:"average"`
	FacedBalls      sql.NullInt64   `json:"faced_balls"`
	StrikeRate      sql.NullFloat64 `json:"strike_rate"`
	ScoreHundreds   sql.NullInt64   `json:"score_hundreds"`
	ScoreFiftys     sql.NullInt64   `json:"score_fiftys"`
	ScoreZeros      sql.NullInt64   `json:"score_zeros"`
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createPlayer,
		arg.Name,
		arg.CareerStartYear,
		arg.CareerEndYear,
		arg.Matches,
		arg.Inns,
		arg.NotOuts,
		arg.Runs,
		arg.HighestScores,
		arg.Average,
		arg.FacedBalls,
		arg.StrikeRate,
		arg.ScoreHundreds,
		arg.ScoreFiftys,
		arg.ScoreZeros,
	)
}

const deleteAllPlayers = `-- name: DeleteAllPlayers :execresult
DELETE FROM players
`

func (q *Queries) DeleteAllPlayers(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteAllPlayers)
}

const getPlayerOfMostRuns = `-- name: GetPlayerOfMostRuns :one
SELECT id, name, career_start_year, career_end_year, matches, inns, not_outs, runs, highest_scores, average, faced_balls, strike_rate, score_hundreds, score_fiftys, score_zeros, created_at, updated_at FROM players
ORDER BY runs DESC LIMIT 1
`

func (q *Queries) GetPlayerOfMostRuns(ctx context.Context) (Player, error) {
	row := q.db.QueryRowContext(ctx, getPlayerOfMostRuns)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CareerStartYear,
		&i.CareerEndYear,
		&i.Matches,
		&i.Inns,
		&i.NotOuts,
		&i.Runs,
		&i.HighestScores,
		&i.Average,
		&i.FacedBalls,
		&i.StrikeRate,
		&i.ScoreHundreds,
		&i.ScoreFiftys,
		&i.ScoreZeros,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPlayerOfMostRunsByCareerEndYear = `-- name: GetPlayerOfMostRunsByCareerEndYear :one
SELECT id, name, career_start_year, career_end_year, matches, inns, not_outs, runs, highest_scores, average, faced_balls, strike_rate, score_hundreds, score_fiftys, score_zeros, created_at, updated_at FROM players
WHERE career_end_year = ?
ORDER BY runs DESC LIMIT 1
`

func (q *Queries) GetPlayerOfMostRunsByCareerEndYear(ctx context.Context, careerEndYear sql.NullInt64) (Player, error) {
	row := q.db.QueryRowContext(ctx, getPlayerOfMostRunsByCareerEndYear, careerEndYear)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CareerStartYear,
		&i.CareerEndYear,
		&i.Matches,
		&i.Inns,
		&i.NotOuts,
		&i.Runs,
		&i.HighestScores,
		&i.Average,
		&i.FacedBalls,
		&i.StrikeRate,
		&i.ScoreHundreds,
		&i.ScoreFiftys,
		&i.ScoreZeros,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPlayersByCareerYear = `-- name: GetPlayersByCareerYear :many
SELECT id, name, career_start_year, career_end_year, matches, inns, not_outs, runs, highest_scores, average, faced_balls, strike_rate, score_hundreds, score_fiftys, score_zeros, created_at, updated_at FROM players
WHERE career_start_year <= ? AND career_end_year >= ?
ORDER By id ASC
`

type GetPlayersByCareerYearParams struct {
	CareerYear sql.NullInt64 `json:"career_year"`
}

func (q *Queries) GetPlayersByCareerYear(ctx context.Context, arg GetPlayersByCareerYearParams) ([]Player, error) {
	rows, err := q.db.QueryContext(ctx, getPlayersByCareerYear, arg.CareerYear, arg.CareerYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Player{}
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CareerStartYear,
			&i.CareerEndYear,
			&i.Matches,
			&i.Inns,
			&i.NotOuts,
			&i.Runs,
			&i.HighestScores,
			&i.Average,
			&i.FacedBalls,
			&i.StrikeRate,
			&i.ScoreHundreds,
			&i.ScoreFiftys,
			&i.ScoreZeros,
			&i.CreatedAt,
			&i.UpdatedAt,
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
