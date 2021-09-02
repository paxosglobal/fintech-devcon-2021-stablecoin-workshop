contract-bindings: SHELL:=/usr/bin/env bash
contract-bindings: contracts/USDK.sol
	# TODO: fill out this function for exercise 1!

start-local: contract-bindings
	docker-compose up

run-backend:
	@go run ./cmd/backend

run-frontend:
	@cd webapp && npm start
