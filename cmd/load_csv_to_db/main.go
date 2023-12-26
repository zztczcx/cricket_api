package main

import (
	"cricket/cmd/load_csv_to_db/batting"
	"flag"
	"fmt"
)

func main() {
	i := flag.String("input", "./db/seeds/ODI_data.csv", "Source file")
	flag.Parse()

	if *i == "" {
		panic("Missing data file")
	}

	loader := batting.NewLoader(i)
	err := loader.Load()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Data imported successfully")
	}
}
