contract-bindings: SHELL:=/usr/bin/env bash
contract-bindings: contracts/USDK.sol
	npm install
	npm run compile
	cat build/contracts/UsdToken.json | jq -c .abi > build/USDK.abi
	cat build/contracts/UsdToken.json | jq  -r .bytecode > build/USDK.bin
	docker run -v $(shell pwd):/sources ethereum/client-go:alltools-v1.10.6 abigen --type USDK \
		--bin="/sources/build/USDK.bin" \
		--abi="/sources/build/USDK.abi" \
		--pkg=contracts --out="/sources/build/USDK.go"

start-local: contract-bindings
	docker-compose up

run-backend:
	@go run ./cmd/backend

run-frontend:
	@cd webapp && npm start

exercise-3:
	@go test ./pkg/server -run TestMint
