#!/bin/bash
docker-compose exec db01 bash -c "echo \"use employees;insert into departments values('${RANDOM:0:4}','${RANDOM:0:4}');\" | mysql -proot"
