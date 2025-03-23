protoc --go_out=../pb/go --go-grpc_out=../pb/go *.proto &

python -m grpc_tools.protoc -I=. --grpc_python_out=../pb/python *.proto &
protoc -I=. --python_out=../pb/python --pyi_out=../pb/python *.proto &

protoc -I=. --cpp_out=../pb/cpp --grpc_out=../pb/cpp --plugin=protoc-gen-grpc="C:\vcpkg\packages\grpc_x64-windows\tools\grpc\grpc_cpp_plugin.exe" *.proto &

protoc -I=. --java_out=../pb/java --grpc-java_out=../pb/java --plugin=protoc-gen-grpc-java="C:\vcpkg\packages\grpc_x64-windows\tools\grpc\grpc_java_plugin-1.68.1.exe" *.proto