@echo on
cd /d ..\..\clients\Demo\
set "cmd=go run grpc_client.go
echo %cmd% | clip
start cmd /k %cmd%