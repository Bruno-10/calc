#### RUN FUNCTIONS
build-api:
	go build -o ./app/services/api/ ./app/services/api 

run-api: build-api
	app/services/api/./api

install-gui:
	yarn --cwd ./app/services/frontend install

run-gui: install-gui
	yarn --cwd ./app/services/frontend run start

run-gui-no-cache: install-gui
	yarn --cwd ./app/services/frontend run start --reset-cache

build-cli:
	go build -o ./app/services/cli ./app/services/cli

run-cli:
	app/services/cli/./cli

#### UTILS

tidy:
	go mod tidy
	go mod vendor

test-race:
	CGO_ENABLED=1 go test -race -count=1 ./...

test-only:
	CGO_ENABLED=0 go test -count=1 ./...

lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

vuln-check:
	govulncheck ./...

test: test-only lint vuln-check

test-race: test-race lint vuln-check

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor
