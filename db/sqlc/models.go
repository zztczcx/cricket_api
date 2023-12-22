// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"database/sql"
	"time"
)

type Player struct {
	ID              int32           `json:"id"`
	Name            string          `json:"name"`
	CareerStartYear sql.NullInt32   `json:"career_start_year"`
	CareerEndYear   sql.NullInt32   `json:"career_end_year"`
	Matches         sql.NullInt32   `json:"matches"`
	Inns            sql.NullInt32   `json:"inns"`
	NotOuts         sql.NullInt32   `json:"not_outs"`
	Runs            sql.NullInt32   `json:"runs"`
	HighestScores   sql.NullInt32   `json:"highest_scores"`
	Average         sql.NullFloat64 `json:"average"`
	FacedBalls      sql.NullInt32   `json:"faced_balls"`
	StrikeRate      sql.NullFloat64 `json:"strike_rate"`
	ScoreHundreds   sql.NullInt32   `json:"score_hundreds"`
	ScoreFiftys     sql.NullInt32   `json:"score_fiftys"`
	ScoreZeros      sql.NullInt32   `json:"score_zeros"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       sql.NullTime    `json:"updated_at"`
}
