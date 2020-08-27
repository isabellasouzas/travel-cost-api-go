#!/usr/bin/env bash
cd ..
cd data
docker start be9a7449361a
docker exec -it be9a7449361a mongo --eval "db.reservations.find()" quickstart
docker exec -it be9a7449361a mongo --eval "db.reservations.findOne({'Nome':'MAUREN LOUISE SGUARIO COELHO DE ANDRADE'})" quickstart

