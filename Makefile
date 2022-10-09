install:
	@sh scripts/install.sh

run:
	@GIN_MODE=debug go run main.go
