#!/usr/bin/env bash
go build
./scripts/run_database.sh
./scripts/db_cleaner.sh
go run main.go