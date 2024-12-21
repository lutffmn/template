# import env
include .env

# initialize go mod
init: 
	@ read -p "Module name: " Name; \
	go mod init $${Name}
	@ rm -rf .git/
	@$(MAKE) dep
	@$(MAKE) env
	git init

# install dependencies
dep:
	@go get -u github.com/jackc/pgx/v5
	@go get -u github.com/spf13/viper
	@go get -u github.com/lutffmn/flow

# generate .env
env:
	@ go run ./helper/env/env.go

# create migration
migrate:
	@ read -p "Migration name: " Name; \
		goose -s -dir $(GOOSE_MIGRATION_DIR) create $${Name} sql

vars := $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir $(GOOSE_MIGRATION_DIR)
# up migration to the most recent version available
up:
	@goose $(vars) up
# up migration by one version
up-one:
	@goose $(vars) up-by-one
# up migration to specific version
up-to:
	@ read -p "Please specify migration version: " Version; \
		goose $(vars) up-to $${Version}
	
# down migration to the most recent version available
down:
	goose $(vars) down 
# down migration to specific version
down-to:
	@ read -p "Please specify migration version: " Version; \
		goose $(vars) down-to $${Version}

# re-run the latest migration
redo:
	@goose $(vars) redo

# rollback all migrations
reset:
	@goose $(vars) reset

# checks migrations status
status:
	@goose $(vars) status

# shows current version of the database
version:
	@goose $(vars) version
	
# build the program
build:
	@go build -o bin/$(BINARY_NAME)

# run the compiled binary
run: build 
	@./bin/$(BINARY_NAME)

# run the program
dev:
	@go run main.go

