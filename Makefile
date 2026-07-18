.PHONY: clear-swag swag-init

clear-swag:
	@echo "clearing docs directory..."
	@rm -rf docs/* 2>/dev/null || true

swag-init: clear-swag
	@echo "generating swagger docs..."
	@swag init -g api/main.go -o docs