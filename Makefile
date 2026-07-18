.PHONY: clear-swag swag-init

clear-swag:
	@echo "clearing docs directory..."
	@rm -rf backend/docs/* 2>/dev/null || true

swag-init: clear-swag
	@echo "generating swagger docs..."
	@swag init -g backend/api/main.go -o backend/docs