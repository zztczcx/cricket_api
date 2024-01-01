package db

import (
	"context"
	"math/rand"

	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CreatePlayer(t *testing.T) {
	s := TxStore(t)

	createRandomPlayer(t, s)
}

func Test_GetPlayersByCareerYear(t *testing.T) {
	s := TxStore(t)

	p1 := createRandomPlayer(t, s)
	p2 := createRandomPlayer(t, s)

	arg := GetPlayersByCareerYearParams{
		CareerYear: ToNullInt64("2010"),
	}

	players, err := s.GetPlayersByCareerYear(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, players)

	require.Equal(t, p1.Name, players[0].Name)
	require.Equal(t, p2.Name, players[1].Name)

	arg = GetPlayersByCareerYearParams{
		CareerYear: ToNullInt64("2040"),
	}

	players, err = s.GetPlayersByCareerYear(context.Background(), arg)
	require.NoError(t, err)
	require.Empty(t, players)
}

func Test_GetPlayerOfMostRuns(t *testing.T) {
	s := TxStore(t)

	p1 := createRandomPlayer(t, s)
	p2 := createRandomPlayer(t, s)
	p3 := createRandomPlayer(t, s)

	players := []CreatePlayerParams{p1, p2, p3}
	sort.SliceStable(players, func(i, j int) bool { return players[i].Runs.Int64 > players[j].Runs.Int64 })

	player, err := s.GetPlayerOfMostRuns(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, player)

	require.Equal(t, player.Name, players[0].Name)
}

func Test_GetPlayerOfMostRunsByCareerEndYear(t *testing.T) {
	s := TxStore(t)

	s.CreatePlayer(context.Background(), CreatePlayerParams{
		Name:          "player_1",
		CareerEndYear: ToNullInt64("2011"),
		Runs:          ToNullInt64("1111")},
	)

	s.CreatePlayer(context.Background(), CreatePlayerParams{
		Name:          "player_2",
		CareerEndYear: ToNullInt64("2011"),
		Runs:          ToNullInt64("1122")},
	)
	s.CreatePlayer(context.Background(), CreatePlayerParams{
		Name:          "player_3",
		CareerEndYear: ToNullInt64("2000"),
		Runs:          ToNullInt64("1111")},
	)

	player, err := s.GetPlayerOfMostRunsByCareerEndYear(context.Background(), ToNullInt64("2011"))
	require.NoError(t, err)
	require.NotEmpty(t, player)

	require.Equal(t, player.Name, "player_2")
}

func createRandomPlayer(t *testing.T, s Store) CreatePlayerParams {
	arg := CreatePlayerParams{
		Name:            "player_" + strconv.Itoa(rand.Intn(1000)),
		CareerStartYear: ToNullInt64("19" + strconv.Itoa(rand.Intn(90)+10)),
		CareerEndYear:   ToNullInt64("20" + strconv.Itoa(rand.Intn(20)+10)),
		Matches:         ToNullInt64(strconv.Itoa(rand.Intn(2000))),
		Inns:            ToNullInt64(strconv.Itoa(rand.Intn(2000))),
		NotOuts:         ToNullInt64(strconv.Itoa(rand.Intn(2000))),
		Runs:            ToNullInt64(strconv.Itoa(rand.Intn(2000))),
		HighestScores:   ToNullInt64(strconv.Itoa(rand.Intn(2000))),
		Average:         ToNullFloat64(strconv.Itoa(rand.Intn(2000))),
		FacedBalls:      ToNullInt64(strconv.Itoa(rand.Intn(2000))),
		StrikeRate:      ToNullFloat64(strconv.Itoa(rand.Intn(2000))),
		ScoreHundreds:   ToNullInt64(strconv.Itoa(rand.Intn(2000))),
		ScoreFiftys:     ToNullInt64(strconv.Itoa(rand.Intn(2000))),
		ScoreZeros:      ToNullInt64(strconv.Itoa(rand.Intn(2000))),
	}

	result, err := s.CreatePlayer(context.Background(), arg)
	require.NoError(t, err)

	id, err := result.LastInsertId()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	return arg
}
