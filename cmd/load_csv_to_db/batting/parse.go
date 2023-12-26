package batting

import (
	db "cricket/db/sqlc"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
)

const (
	parserCount = 5
)

type player db.CreatePlayerParams

func (l *loader) startParser(dataSource <-chan string, done chan<- struct{}) {
	var wg sync.WaitGroup
	wg.Add(parserCount)

	for i := 0; i < parserCount; i++ {
		playerChan := make(chan player)
		go parse(dataSource, playerChan)
		go l.Insert(playerChan, &wg)
	}

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()
}

func parse(dataSource <-chan string, playerChan chan<- player) {
	for d := range dataSource {
		player, err := parseData(d)
		if err != nil {
			log.Println(err)
		} else {
			playerChan <- player
		}
	}
	close(playerChan)
}

func parseData(line string) (p player, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error parsing data: %v", r)
		}
	}()

	fields := strings.Split(line, ",")
	careerYear := fields[2]
	yearParts := strings.Split(careerYear, "-")

	p = player{
		Name:            fields[1],
		CareerStartYear: toNullInt64(yearParts[0]),
		CareerEndYear:   toNullInt64(yearParts[1]),
		Matches:         toNullInt64(fields[3]),
		Inns:            toNullInt64(fields[4]),
		NotOuts:         toNullInt64(fields[5]),
		Runs:            toNullInt64(fields[6]),
		HighestScores:   toNullInt64(fields[7]),
		Average:         toNullFloat64(fields[8]),
		FacedBalls:      toNullInt64(fields[9]),
		StrikeRate:      toNullFloat64(fields[10]),
		ScoreHundreds:   toNullInt64(fields[11]),
		ScoreFiftys:     toNullInt64(fields[12]),
		ScoreZeros:      toNullInt64(fields[13]),
	}
	return
}

func toNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64: int64(i), Valid: err == nil}
}

func toNullFloat64(s string) sql.NullFloat64 {
	s = sanitize(s)
	f, err := strconv.ParseFloat(s, 64)
	return sql.NullFloat64{Float64: f, Valid: err == nil}
}

func sanitize(s string) string {
	return strings.Replace(s, "*", "", -1)
}
