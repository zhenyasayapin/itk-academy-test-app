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