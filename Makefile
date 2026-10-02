WALLET_BINARY_BINARY=walletApp

up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

up_build: build_wallet
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

build_wallet:
	@echo "Building wallet binary..."
	cd ./wallet-service && env GOOS=linux CGO_ENABLED=0 go build -o ${WALLET_BINARY_BINARY} ./cmd/api
	@echo "Done!"

migrate:
	@echo "Applying migrations..."	
	cd ./wallet-service && soda migrate
	@echo "Done!"
fixtures:
	docker compose exec -T postgres psql -U postgres -d itk-academy-test-app < ./wallet-service/sql/wallets.sql

benchmark:
	@wrk -t10 -c100 -d10s --latency -s ./post_wallet.lua http://localhost:8080/api/v1/wallet

test:
	cd ./wallet-service/cmd/api && go test -v
