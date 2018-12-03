#!/usr/bin/env bash

pkill portfold

echo "Running go fmt..."
go fmt ./...

echo "Running unit tests..."
go test ./... || exit

echo "Building application and starting..."
go build && ./portfold > /dev/null 2> /dev/null &

echo "Waiting for server to startup..."
while :
do
    # -s is silent mode and -m sets the connection timeout to 1 second
    curl -s -m 1  "http://localhost:8080/health-check/" 2>&1 | grep "health check"
    # $? is the exit value of the last command and if it passes break out of the loop
    [[ $? ]] && break
done

echo "Running integration tests..."

GREEN='\033[1;32m'
RED='\033[0;31m'
NC='\033[0m' # no colour

if go test -tags=integration ./... ; then
	echo -e "${GREEN}--------- ALL TESTS PASSED ---------"
else
	exit
fi

echo -e "${NC}"
pkill portfold