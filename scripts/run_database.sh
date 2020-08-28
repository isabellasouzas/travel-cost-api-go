#!/usr/bin/env bash
docker start be9a7449361a
docker cp ./data/data.json be9a7449361a:/tmp/data.json
docker exec be9a7449361a mongoimport -d quickstart -c reservations --file /tmp/data.json --jsonArray
echo "Database is running..."

