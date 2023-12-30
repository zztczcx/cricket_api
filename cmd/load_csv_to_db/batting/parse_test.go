package batting

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseData(t *testing.T) {
	line := "0,SR Tendulkar (INDIA),1989-2012,463,452,41,18426,200*,44.83,21367,86.23,49,96,20,"
	p, err := parseData(line)

	require.NoError(t, err)
	require.Equal(t, "SR Tendulkar (INDIA)", p.Name)
	require.Equal(t, sql.NullInt64{Int64: 1989, Valid: true}, p.CareerStartYear)
	require.Equal(t, sql.NullInt64{Int64: 2012, Valid: true}, p.CareerEndYear)
	require.Equal(t, sql.NullInt64{Int64: 200, Valid: true}, p.HighestScores)
	require.Equal(t, sql.NullFloat64{Float64: 44.83, Valid: true}, p.Average)

	//HighestScores empty
	line2 := "0,SR Tendulkar (INDIA),1989-2012,463,452,41,18426,,44.83,21367,86.23,49,96,20,"
	p, err = parseData(line2)

	require.NoError(t, err)
	require.Equal(t, sql.NullInt64{Int64: 0, Valid: false}, p.HighestScores)

	//career span empty
	line3 := "0,SR Tendulkar (INDIA),,463,452,41,18426,,44.83,21367,86.23,49,96,20,"
	_, err = parseData(line3)

	require.ErrorContains(t, err, "error parsing data")
}
