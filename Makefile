include env_make

migrate_up:
	migrate -database $(POSTGRES_URI) -path migrations up

migrate_down:
	migrate -database $(POSTGRES_URI) -path migrations down

mocks:
	mockgen -source=internal/user/usecase/interface.go -destination=internal/user/usecase/mock/user.go -package=mock

test:
	go test -v -coverprofile cover.out -tags testing ./...
	go tool cover -html=cover.out -o coverage.html
