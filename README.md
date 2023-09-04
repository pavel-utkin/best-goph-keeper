Выпускной проект (№2) Яндекс Практикума "Менеджер паролей GophKeeper"

[![Golang](https://img.shields.io/badge/-Go-0000FF?style=flat-square&logo=Go)](https://go.dev)
[![gRPC](https://img.shields.io/badge/-gRPC-0000FF?style=flat-square&logo=gRPC)](https://grpc.io)
[![protobuf](https://img.shields.io/badge/-protobuf-0000FF?style=flat-square&logo=protobuf)](https://protobuf.dev)
[![PostgreSQL](https://img.shields.io/badge/-PostgreSQL-0000FF?style=flat-square&logo=PostgreSQL)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/-Docker-0000FF?style=flat-square&logo=Docker)](https://docker.com/)

- [Техническое задание](SPECS.md)


Параметры запуска сервера

1) AddressGRPC - grpcserver address (ex. localhost:8080) `(flag -g)`
2) AddressREST - rest server address (ex. localhost:8088) `(flag -r)`
3) DATABASE_DSN - database address (ex. DEBUG) `(flag -d)`
5) DATA_FOLDER - folder to save user data `(flag -f)`

Параметры запуска клинта

1) GRPC - grpc server address (ex. localhost:8080) `(flag -g)`
3) DATA_FOLDER - folder to save user data`(flag -f)`


Запуск

 1) Открыть терминал и ввести команду `make dev-up`
 2) Запуск сервера `./bin/server -g "localhost:8080" -r "localhost:8088" -d "host=localhost port=5432 user=postgres password=password dbname=gophkeeper sslmode=disable" -f "./data/server_keeper"`
 3) Запуск клиента `./bin/client -g "localhost:8080" -f "./data/client_keeper"`
