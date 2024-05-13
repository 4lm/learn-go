dev:
	@echo "Executing make:dev..."
	@echo "Checking if MySQL is running..."
	@pgrep mysql > /dev/null || (echo "MySQL is not running..." && mysql.server start)
	@mysqladmin ping --silent || (echo "Exiting. MySQL Connection was not established."; exit 1)
	@echo "Starting Go web server..."
	@. ./.env && go run ./cmd/web -dsn=$$PASTY_DSN
