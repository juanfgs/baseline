#!/usr/bin/sh
#Replace the tokens for the appropiate values

migrate -url=mysql://MYSQL_USER:MYSQL_PASS@tcp\(MYSQL_HOST:MYSQL_PORT\)/MYSQL_DBNAME -path=./migrations $1 $2

