```sql
CREATE USER 'go'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'go'@'%';
flush privileges;

CREATE SCHEMA `go_db`;
```

`MYSQL_ROOT_PASSWORD` env variable needs to be set in Docker image container for MySQL8

0 row(s) affected, 1 warning(s): 1285 MySQL is started in --skip-name-resolve mode; you must restart it without this switch for this grant to work
