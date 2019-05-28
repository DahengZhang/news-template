docker run -d --rm --name new-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=000000 mysql

docker exec -it new-mysql env LANG=C.UTF-8 sh

mysql -uroot -p

USE mysql;

CREATE USER "dahengzhang"@"%" IDENTIFIED BY "000000";

GRANT ALL ON *.* TO "dahengzhang"@"%" WITH GRANT OPTION;

FLUSH PRIVILEGES;

CREATE DATABASE webIM;
