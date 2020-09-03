#!/usr/bin/env bash
./run_database.sh
./db_cleaner.sh
go build
go run main.go