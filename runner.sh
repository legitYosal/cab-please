#! /bin/bash

docker-compose down
docker-compose build
docker-compose up -d --scale decreaser=4