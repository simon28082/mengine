.PHONY:install
install:
	@echo "install google wire"
	go install github.com/google/wire/cmd/wire@latest

.PHONY:build
build:
	@echo "build wire"
	wire ./...

.PHONY:run
run:
	go run -tags '!wireinject' cmd/wire_gen.go