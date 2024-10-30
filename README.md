# go-ecommerce-api

run mysql 

```
docker run --name mysql -e MYSQL_ROOT_PASSWORD=mypassword -d -p 3306:3306 -v mysql-data:/var/lib/mysql mysql:9.1.0
mysql -u root -p -h 127.0.0.1 -P 3306
CREATE DATABASE ecommerce;
```