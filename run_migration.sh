#!/usr/bin/sh

migrate -url=mysql://root:fusion87@tcp\(localhost:3306\)/cv -path=./migrations $1 $2

