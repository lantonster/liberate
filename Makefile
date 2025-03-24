.PHONY: run
run:
	go run cmd/main.go

.PHONY: wire
wire:
	wire gen cmd/wire/wire.go

.PHONY: orm
orm:
	go run cmd/orm/main.go