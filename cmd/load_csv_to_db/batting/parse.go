package batting

import (
	db "cricket/db/sqlc"
	"fmt"
	"strings"
	"sync"
)

const (
	parserCount = 5
)

type player db.CreatePlayerParams

func (l *loader) startParser(dataSource <-chan string, done chan<- struct{}, errors chan<- error) {
	var wg sync.WaitGroup
	wg.Add(parserCount)

	for i := 0; i < parserCount; i++ {
		playerChan := make(chan player)
		go parse(dataSource, playerChan, errors)
		go l.Insert(playerChan, &wg, errors)
	}

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()
}

func parse(dataSource <-chan string, playerChan chan<- player, errors chan<- error) {
	for d := range dataSource {
		player, err := parseData(d)
		if err != nil {
                        errors <- err
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
		CareerStartYear: db.ToNullInt64(yearParts[0]),
		CareerEndYear:   db.ToNullInt64(yearParts[1]),
		Matches:         db.ToNullInt64(fields[3]),
		Inns:            db.ToNullInt64(fields[4]),
		NotOuts:         db.ToNullInt64(fields[5]),
		Runs:            db.ToNullInt64(fields[6]),
		HighestScores:   db.ToNullInt64(sanitize(fields[7])),
		Average:         db.ToNullFloat64(sanitize(fields[8])),
		FacedBalls:      db.ToNullInt64(fields[9]),
		StrikeRate:      db.ToNullFloat64(sanitize(fields[10])),
		ScoreHundreds:   db.ToNullInt64(fields[11]),
		ScoreFiftys:     db.ToNullInt64(fields[12]),
		ScoreZeros:      db.ToNullInt64(fields[13]),
	}
	return
}

func sanitize(s string) string {
	return strings.Replace(s, "*", "", -1)
}
