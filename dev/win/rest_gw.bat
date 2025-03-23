@echo off
cd /d ..\..\rest-gateway\

set port=localhost:5001
set grpcgw=localhost:5000
start cmd /k "color 0C & go run ."