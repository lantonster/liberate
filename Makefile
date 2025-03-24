.PHONY: start start-web run

start:
	go run cmd/main.go

start-web:
	cd web && npm run dev

run:
	make start &
	make start-web