compose-build:
	@docker compose build

compose-up:
	@docker compose up -d

compose-down:
	@docker compose down

compose-restart: compose-down compose-build compose-up


