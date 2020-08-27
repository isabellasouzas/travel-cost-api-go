#!/usr/bin/env bash
go build
./db_script/run.sh
./db_script/db_cleaner.sh
go run main.go