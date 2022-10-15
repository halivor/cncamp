#!/bin/bash

cp -rf ../module2 .
docker build . -t cncamp/httpserver:v1.0.0
docker run -d cncamp/httpserver:v1.0.0

rm -rf ./module2
PID=`ps -ef | grep "exec" | grep server | grep -v grep | awk '{print $2}'`
echo $PID
nsenter -t $PID -n ip addr
