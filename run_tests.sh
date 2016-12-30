#!/bin/sh

export APP_ENV=testing
CDIR=$(pwd)

cd $CDIR/handlers
go test
cd $CDIR/models/
go test
