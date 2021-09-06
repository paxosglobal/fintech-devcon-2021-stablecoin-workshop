# prerequisites
# install golang https://golang.org/doc/install
# install docker https://docs.docker.com/get-docker/
# install node https://nodejs.org/en/download/

# for a quicker option on mac, you can use the homebrew based helper: `make bootstrap-homebrew`
bootstrap-homebrew:
	which brew || /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
	which go || brew install go
	which docker || brew cask install docker
	which node || brew install node

contract-bindings: SHELL:=/usr/bin/env bash
contract-bindings: contracts/USDK.sol
	npm install
	npm run compile
	cat build/contracts/UsdToken.json | ./node_modules/node-jq/bin/jq -c .abi > build/USDK.abi
	cat build/contracts/UsdToken.json | ./node_modules/node-jq/bin/jq -r .bytecode > build/USDK.bin
	docker run -v $(shell pwd):/sources ethereum/client-go:alltools-v1.10.6 abigen --type USDK \
		--bin="/sources/build/USDK.bin" \
		--abi="/sources/build/USDK.abi" \
		--pkg=contracts --out="/sources/build/USDK.go"

start-local: contract-bindings
	docker-compose up

run-backend:
	@go run ./cmd/backend

run-frontend:
	@cd webapp && npm install && npm start

exercise-3:
	@go test ./pkg/server -run TestMint
