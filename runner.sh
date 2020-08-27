#!/usr/bin/env bash
go build
./db_connection_script/run.sh
go run main.go