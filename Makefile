migration-up:
	@cd migration/ && goose postgres "user=fitMachine password=fitMachine sslmode=disable host=127.0.0.1 port=5432 database=fitMachine" up

migration-down:
	@cd migration/ && goose postgres "user=fitMachine password=fitMachine sslmode=disable host=127.0.0.1 port=5432 database=fitMachine" down