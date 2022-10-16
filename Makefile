install:
	@sh scripts/install.sh

run:
	@GIN_MODE=debug go run main.go

doc:
	@redocly preview-docs docs/openapi.yaml
