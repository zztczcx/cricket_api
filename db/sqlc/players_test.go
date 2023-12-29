package db

import (
	"context"
	"math/rand"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPlayer(t *testing.T) CreatePlayerParams {
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

	player, err := testStore.CreatePlayer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, player)

	return arg
}

func clearTable(t *testing.T) {
	testStore.DeleteAllPlayers(context.Background())
}

func Test_CreatePlayer(t *testing.T) {
	defer clearTable(t)
	createRandomPlayer(t)
}

func Test_GetPlayersByCareerYear(t *testing.T) {
	defer clearTable(t)

	p1 := createRandomPlayer(t)
	p2 := createRandomPlayer(t)

	arg := GetPlayersByCareerYearParams{
		CareerYear: ToNullInt64("2010"),
	}

	players, err := testStore.GetPlayersByCareerYear(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, players)

	require.Equal(t, p1.Name, players[0].Name)
	require.Equal(t, p2.Name, players[1].Name)

	arg = GetPlayersByCareerYearParams{
		CareerYear: ToNullInt64("2040"),
	}

	players, err = testStore.GetPlayersByCareerYear(context.Background(), arg)
	require.NoError(t, err)
	require.Empty(t, players)
}

func Test_GetPlayerOfMostRuns(t *testing.T) {
	defer clearTable(t)

	p1 := createRandomPlayer(t)
	p2 := createRandomPlayer(t)
	p3 := createRandomPlayer(t)

	players := []CreatePlayerParams{p1, p2, p3}
	sort.SliceStable(players, func(i, j int) bool { return players[i].Runs.Int64 > players[j].Runs.Int64 })

	player, err := testStore.GetPlayerOfMostRuns(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, player)

	require.Equal(t, player.Name, players[0].Name)
}

func Test_GetPlayerOfMostRunsByCareerEndYear(t *testing.T) {
	defer clearTable(t)

	testStore.CreatePlayer(context.Background(), CreatePlayerParams{
		Name:          "player_1",
		CareerEndYear: ToNullInt64("2011"),
		Runs:          ToNullInt64("1111")},
	)

	testStore.CreatePlayer(context.Background(), CreatePlayerParams{
		Name:          "player_2",
		CareerEndYear: ToNullInt64("2011"),
		Runs:          ToNullInt64("1122")},
	)
	testStore.CreatePlayer(context.Background(), CreatePlayerParams{
		Name:          "player_3",
		CareerEndYear: ToNullInt64("2000"),
		Runs:          ToNullInt64("1111")},
	)

	player, err := testStore.GetPlayerOfMostRunsByCareerEndYear(context.Background(), ToNullInt64("2011"))
	require.NoError(t, err)
	require.NotEmpty(t, player)

	require.Equal(t, player.Name, "player_2")
}
