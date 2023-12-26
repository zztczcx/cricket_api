package batting

import (
	"bufio"
	"log"
	"os"
)

func (l *loader) produce() <-chan string {
	dataSource := make(chan string)

	go func() {
		f, err := os.Open(*l.dataFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		scanner.Scan() // skip header

		for scanner.Scan() {
			dataSource <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		close(dataSource)

	}()

	return dataSource
}
