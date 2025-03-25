# 运行项目
.PHONY: run
run:
	go run cmd/main.go

# 生成 wire 代码
.PHONY: wire
wire:
	wire gen cmd/wire/wire.go

# 生成 orm 代码
.PHONY: orm
orm:
	go run cmd/orm/main.go

# 生成 swagger 文档
.PHONY: swag
swag:
	swag fmt && swag init -g internal/router/router.go -o docs/api --parseDependency
