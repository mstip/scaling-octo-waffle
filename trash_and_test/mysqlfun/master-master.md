# master - master no downtime

1. add db03 as slave
2. set db01 as slave from db03 (master master)
3. change all writes to db03
4. lock db01 to have definitly only write on db03
5. stop slave on db01 and db03
6. release lock and drop database on db01

ensure
- binlog and server name enabled everywhere / permissions to do so
- replication user setup or permission to do so
- permissions slave change master commands
- drop database permissions
- lock table permission
- check that db02 is only read or best not used !
- check that master slave stuff is not somewhere hidden in conf ask how it is setup


NOTE: after restarting to clear
```
docker-compose down
docker volume prune
```

## initial setup
- master - slave replica between db01 and db02
- db01 has 2 databases
- db03 db04 empty

## Initial setup master-slave db 01 db 02
```
docker cp test_db/ db01:/root
docker-compose exec db01 bash -c "cd /root/test_db;mysql -proot < employees.sql"
docker-compose exec db01 bash -c "cd /root/test_db/sakila;mysql -proot < sakila-mv-schema.sql"
docker-compose exec db01 bash -c "cd /root/test_db/sakila;mysql -proot < sakila-mv-data.sql"
docker cp master-master/ db01:/root
docker cp master-master/ db02:/root
docker-compose exec db01 bash -c "cd /root/master-master;mysql -proot < db01.sql"

docker-compose exec db01 bash -c "mysqldump -proot --skip-lock-tables --single-transaction --flush-logs --master-data=2 -A > /root/mysqldump.sql"
docker cp db01:/root/mysqldump.sql .
docker cp mysqldump.sql db02:/root/mysqldump.sql
rm mysqldump.sql
docker-compose exec db02 bash -c "cd /root;mysql -proot < mysqldump.sql"
docker-compose exec db02 bash -c "cd /root;head mysqldump.sql -n80 | grep "MASTER_LOG_POS"" // should be -- CHANGE MASTER TO MASTER_LOG_FILE='mysql-bin.000007', MASTER_LOG_POS=154;
docker-compose exec db02 bash -c "cd /root/master-master;mysql -proot < db02.sql"

```

## start watchers readers
```
watch "docker-compose exec db01 bash -c 'echo \"use employees;select count(*) from departments;\" | mysql -proot '"
watch "docker-compose exec db02 bash -c 'echo \"use employees;select count(*) from departments;\" | mysql -proot '"
watch "docker-compose exec db01 bash -c 'echo \"use sakila;select count(*) from category;\" | mysql -proot '"
watch "docker-compose exec db02 bash -c 'echo \"use sakila;select count(*) from category;\" | mysql -proot '"
```

## start watchers writer
```
watch ./master-master/write_to_db01_employees_deps.sh
watch ./master-master/write_to_db01_sakila_category.sh 
```

## add db03 as slave
```
docker-compose exec db01 bash -c "mysqldump -proot --skip-lock-tables --single-transaction --flush-logs --master-data=2 -A > /root/mysqldump.sql"
docker cp db01:/root/mysqldump.sql .
docker cp mysqldump.sql db03:/root/mysqldump.sql
rm mysqldump.sql
docker-compose exec db03 bash -c "cd /root;mysql -proot < mysqldump.sql"
docker-compose exec db03 bash -c "cd /root;head mysqldump.sql -n80 | grep "MASTER_LOG_POS""
# !!!DONT FORGET TO CHANGE THE VALUES HERE!!!
docker-compose exec db03 bash -c "echo \"CHANGE MASTER TO MASTER_HOST='db01',MASTER_USER='replicant',MASTER_PASSWORD='password', MASTER_LOG_FILE='mysql-bin.000009', MASTER_LOG_POS=154;START SLAVE;\" | mysql -proot"
```

## start watchers readers
```
watch "docker-compose exec db03 bash -c 'echo \"use employees;select count(*) from departments;\" | mysql -proot '"
watch "docker-compose exec db03 bash -c 'echo \"use sakila;select count(*) from category;\" | mysql -proot '"


watch "docker-compose exec db01 bash -c 'echo \"use sakila;select category_id from category order by category_id desc limit 1;\" | mysql -proot '"
watch "docker-compose exec db03 bash -c 'echo \"use sakila;select category_id from category order by category_id desc limit 1;\" | mysql -proot '"



```

## setup master master
```
docker-compose exec db03 bash -c "echo \"CREATE USER 'repli'@'%';GRANT REPLICATION SLAVE ON *.* TO 'repli'@'%' IDENTIFIED BY 'password';FLUSH PRIVILEGES;\" | mysql -proot"
docker-compose exec db03 bash -c "echo \"SHOW MASTER STATUS\" | mysql -proot"
docker-compose exec db01 bash -c "echo \"slave stop; CHANGE MASTER TO MASTER_HOST='db03',MASTER_USER='repli',MASTER_PASSWORD='password', MASTER_LOG_FILE='mysql-bin.000001', MASTER_LOG_POS=70943935;START SLAVE;\" | mysql -proot"
```

## checkout !
NOTE ! in this moment db02 is garbage because it only replicates from db01
```
docker-compose exec db03 bash -c "echo \"use employees;insert into departments values('xxx','xxx');\" | mysql -proot"
docker-compose exec db03 bash -c "echo \"use employees;select * from departments;\" | mysql -proot"
docker-compose exec db01 bash -c "echo \"use employees;select * from departments;\" | mysql -proot"

docker-compose exec db01 bash -c "echo \"use employees;insert into departments values('yyy','yyy');\" | mysql -proot"
docker-compose exec db01 bash -c "echo \"use employees;select * from departments;\" | mysql -proot"
docker-compose exec db03 bash -c "echo \"use employees;select * from departments;\" | mysql -proot"
```

## now change all writes to db03 and lock tables on db01 all writes to db01 should fail
```
lock interactive
lock table departments write;
UNLOCK TABLES;
```

## you could now only write to db03 so make it to only master
```
# on db03
stop slave;
RESET SLAVE ALL;
# on db01
stop slave;
RESET SLAVE ALL;
```