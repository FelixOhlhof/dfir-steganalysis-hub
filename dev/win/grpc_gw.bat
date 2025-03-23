@echo off
cd /d ..\..\grpc-gateway\

set port=localhost:5000
set services=localhost:30124;localhost:30125;localhost:30126;localhost:30127

start cmd /k "color 0F & go run ."