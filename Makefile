# ===== CONFIG =====
DB_URL=mysql://root:alif13579@tcp(127.0.0.1:3306)/go_sqlc_db
MIGRATION_PATH=migrations

# ===== MIGRATION =====
migrate-up:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" down 1

migrate-force:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" force $(v)

migrate-version:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" version

# ===== SQLC =====
sqlc:
	sqlc generate

# ===== RUN APP =====
run:
	go run cmd/app/main.go

# ===== BUILD =====
build:
	go build -o bin/app cmd/app/main.go

# ===== CLEAN =====
clean:
	rm -rf bin/