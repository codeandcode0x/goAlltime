# Gin scaffold
This scaffold based [gin](https://github.com/gin-gonic/gin) web framework, you can use it create web project quickly.

# Quick start

```sh
#run mariadb
docker run -p 127.0.0.1:3306:3306  --name mariadb -e MARIADB_ROOT_PASSWORD=root123 -d mariadb:10.2.38

# create database
CREATE SCHEMA `gin_scaffold` DEFAULT CHARACTER SET utf8mb4 ;

# run project
go run main.go

```

# gRPC

```sh
GO111MODULE=off protoc -I helloworld/ helloworld/pb/helloworld.proto --go_out=plugins=grpc:helloworld

protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto

protoc -I grpc/protos/jobs/ grpc/protos/job/*proto --go_out=plugins=grpc:grpc/protos/job
```