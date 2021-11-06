#!/bin/bash
docker-compose exec db01 bash -c "echo \"use sakila;insert into category(name) values('${RANDOM}');\" | mysql -proot"
