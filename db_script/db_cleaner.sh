#!/usr/bin/env bash
cd ..
cd data
docker start be9a7449361a
docker exec -it be9a7449361a mongo --eval "db.reservations.find().forEach(\
function(doc){\
doc['Periodo - Data de inicio']=new Date(doc['Periodo - Data de inicio']);
doc['Periodo - Data de fim']=new Date(doc['Periodo - Data de fim']);
db.reservations.save(doc);})" quickstart
docker exec -it be9a7449361a mongo --eval "db.reservations.find().limit(1).pretty()" quickstart
