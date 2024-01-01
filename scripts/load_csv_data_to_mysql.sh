#!/bin/bash

if [ -z "$1" ]; then
  go run cmd/load_csv_to_db/main.go
else
  go run cmd/load_csv_to_db/main.go -input=$1
fi

