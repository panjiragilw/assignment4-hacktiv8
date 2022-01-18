server:
	nodemon --exec go run main.go --signal SIGTERM

swag:
	swag init --parseDependency --parseInternal

swag-server:
	swag init --parseDependency --parseInternal
	nodemon --exec go run main.go --signal SIGTERM

.PHONY: sqlc server


