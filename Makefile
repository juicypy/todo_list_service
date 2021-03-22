deps-verify:
	@go mod tidy
	@go mod verify

run-server:
	go run cmd/server/main.go

install-deps:
	@GO111MODULE=off go get -u github.com/golang-migrate/migrate/cmd/migrate

migration-new: #Create new migration
	migrate create -ext sql -dir migrations $(name)

migrate-up:
	@go run cmd/migrate/main.go -direction up

migrate-down:
	@go run cmd/migrate/main.go -direction down

migrate-to-version:
	@go run cmd/migrate/main.go -version $(version)
