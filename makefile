## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up: 
	@echo Stopping docker images (if running...)
	docker compose down
	@echo create all container and run it.
	docker compose up --build -d
	@echo success crete and run docker container

## down: stop docker compose
down:
	@echo Stopping docker compose...
	docker compose down
	@echo Done!