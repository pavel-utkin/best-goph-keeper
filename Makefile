.PHONY: up
up:
	docker-compose up

.PHONY: dev-up
dev-up:
	docker-compose -f docker-compose.dev.yaml up

.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -seq -dir ./migrations name

.PHONY: migrate-up
migrate-up:
	migrate -source file://./migrations -database 'postgres://user:password@localhost:5432/gophkeeper?sslmode=disable' up 5

.PHONY: migrate-down
migrate-down:
	migrate -source file://./migrations -database 'postgres://user:password@localhost:5432/gophkeeper?sslmode=disable' down 5

.PHONY: migrate-up-dev
migrate-up:
	migrate -source file://./migrations -database 'postgres://user:password@localhost:5433/gophkeeper_dev?sslmode=disable' up 5

.PHONY: migrate-down-dev
migrate-down:
	migrate -source file://./migrations -database 'postgres://user:password@localhost:5433/gophkeeper_dev?sslmode=disable' down 5


#docs url - http://localhost:6060/pkg/?m=all
.PHONY: run_godoc
run_godoc:
	godoc -http=:6060

.PHONY: gen_protoc
gen_protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
  	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	internal/server/proto/gophkeeper.proto

.PHONY: build_app
build_app:
	go build -o bin/client cmd/client/main.go
	go build -o bin/server cmd/server/main.go

#GOOS=linux GOARCH=amd64  go build -o bin/linux64_client cmd/client/main.go
#GOOS=linux GOARCH=amd64  go build -o bin/linux64_server cmd/server/main.go
#GOOS=windows GOARCH=amd64  go build -o bin/win64_client.exe cmd/client/main.go
#GOOS=windows GOARCH=amd64  go build -o bin/win64_server.exe cmd/server/main.go
#GOOS=darwin GOARCH=amd64 go build -o bin/mac64_client cmd/client/main.go
#GOOS=darwin GOARCH=amd64 go build -o bin/mac64_server cmd/server/main.go