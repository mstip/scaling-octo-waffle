CREATE USER 'replicant'@'%';GRANT REPLICATION SLAVE ON *.* TO 'replicant'@'%' IDENTIFIED BY 'password';FLUSH PRIVILEGES;