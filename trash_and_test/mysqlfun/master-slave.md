# master - slave short downtime
1. add db03 as slave
2. lock all tables on db01
3. stop slave on db03
4. change all writes to db03
5. release lock on db01 and drop table

cronjob
stop worker
add die
inform bi for new slave
ask sysadmins mysqldump

1. add db03 as slave and wait till it cathes up
2. stop workers on web01 - web 02 has now running workers
3. add die(); to index.php of psp on web01 and web02
4. lock tables on db01 but not migrations
5. stop slave on db03
6. release with new credentials to db03
7. check workers running and die is removed
8. release lock on db01 and drop tgp_psp


lock table application write;                     
lock table authentication_nonce write;            
lock table config write;                          
lock table game write;                            
lock table game_application write;                
lock table hmac_authentication write;             
lock table migration_versions write;              
lock table payment write;                         
lock table payment_app_data write;                
lock table payment_customer_data write;           
lock table payment_event write;                   
lock table payment_incoming_provider_event write; 
lock table payment_method write;                  
lock table payment_method_country_status write;   
lock table payment_method_game_country write;     
lock table payment_method_limits write;           
lock table payment_method_limits_country  write;  
lock table payment_provider write;                
lock table provider_adyen_sca_api_log write;      
lock table provider_config write;                 
lock table provider_game_config  write;           
lock table transaction_log  write;                





TODO: 
- how to check psp is running 
- worker control stop
- new credentials
- access to db03


DB SERVER 
- ssh tgdev@db01.prod.paynet.solutions.cgn.travian.info 
- ssh tgdev@db02.prod.paynet.solutions.cgn.travian.info 
- ssh tgdev@db03.prod.paynet.solutions.cgn.travian.info 
- ssh tgdev@db04.prod.paynet.solutions.cgn.travian.info 

WEB SERVER
- ssh m.stipcevic@web01.prod.paynet.solutions.cgn.travian.info
- ssh m.stipcevic@web02.prod.paynet.solutions.cgn.travian.info

BAMBOO PROD
https://bamboo.traviangames.com/deploy/config/configureDeploymentProject.action?id=117669889&environmentId=123600903&returnUrl=/deploy/viewAllDeploymentProjects.action

SEE WORKERS
ps aux  grep tgp-psp

stop cronjobs
crontab -e and comment out all tgp-psp

TODO: Stop workers
./console worker:control stop 

PATH TO INDEX PHP
/var/www/traviangames.com/tgp-psp.traviangames.com/web/index.php




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
docker-compose exec db02 bash -c "cd /root;head mysqldump.sql -n80  grep "MASTER_LOG_POS"" // should be -- CHANGE MASTER TO MASTER_LOG_FILE='mysql-bin.000007', MASTER_LOG_POS=154;
docker-compose exec db02 bash -c "echo \"stop slave;CHANGE MASTER TO MASTER_HOST='db01',MASTER_USER='replicant',MASTER_PASSWORD='password', MASTER_LOG_FILE='mysql-bin.000005', MASTER_LOG_POS=154;START SLAVE;\"  mysql -proot"

```

## check 
```
docker-compose exec db01 bash -c "echo \"use employees;insert into departments values('yyy','yyy');\"  mysql -proot"
docker-compose exec db01 bash -c "echo \"use employees;select * from departments;\"  mysql -proot"
docker-compose exec db02 bash -c "echo \"use employees;select * from departments;\"  mysql -proot"
```

## add db03 as slave
```
docker-compose exec db01 bash -c "mysqldump -proot --skip-lock-tables --single-transaction --flush-logs --master-data=2 -A > /root/mysqldump.sql"
docker cp db01:/root/mysqldump.sql .
docker cp mysqldump.sql db03:/root/mysqldump.sql
rm mysqldump.sql
docker-compose exec db03 bash -c "cd /root;mysql -proot < mysqldump.sql"
docker-compose exec db03 bash -c "cd /root;head mysqldump.sql -n80  grep "MASTER_LOG_POS""
# !!!DONT FORGET TO CHANGE THE VALUES HERE!!!
docker-compose exec db03 bash -c "echo \"CHANGE MASTER TO MASTER_HOST='db01',MASTER_USER='replicant',MASTER_PASSWORD='password', MASTER_LOG_FILE='mysql-bin.000009', MASTER_LOG_POS=154;START SLAVE;\"  mysql -proot"
```

## check 
```
docker-compose exec db01 bash -c "echo \"use employees;insert into departments values('xxx','xxx');\"  mysql -proot"
docker-compose exec db01 bash -c "echo \"use employees;select * from departments;\"  mysql -proot"
docker-compose exec db02 bash -c "echo \"use employees;select * from departments;\"  mysql -proot"
docker-compose exec db03 bash -c "echo \"use employees;select * from departments;\"  mysql -proot"
```

## lock tables on db01 (interactive)
```
use employees;
lock table departments write;
# writes hang now
docker-compose exec db01 bash -c "echo \"use employees;insert into departments values('aaa','aaa');\"  mysql -proot"
#UNLOCK TABLES;
```

## change writes to db03 (example)
```
docker-compose exec db03 bash -c "echo \"use employees;insert into departments values('aaa','aaa');\"  mysql -proot"
# this hangs
docker-compose exec db01 bash -c "echo \"use employees;select * from departments;\"  mysql -proot"
# this hasnt the value
docker-compose exec db02 bash -c "echo \"use employees;select * from departments;\"  mysql -proot"
# has the value
docker-compose exec db03 bash -c "echo \"use employees;select * from departments;\"  mysql -proot"
```

## stop slave on db03
```
docker-compose exec db03 bash -c "echo \"stop slave; reset slave all;\"  mysql -proot"
```
### drop table release lock db01
```
drop table employees;
unlock tables;
```
