## default arguments
config = "./config.yaml"

## test: Run test and enforce go coverage
test:
	$(eval OUT = $(shell go test ./... -coverprofile cp.out ))
	$(eval COVERAGE_CURRENT = $(shell go tool cover -func=cp.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}' ))

	@echo "tests completed!";
	@echo "coverage passed threshold: $(COVERAGE_CURRENT)%";

## coverage: Show go coverage
coverage: test
	@echo "coverage details:";
	@go tool cover -func=cp.out

## coverage-web: Show go coverage in web
coverage-web: test
	@go tool cover -html=cp.out -o coverage.html

## install: Install module requirement applications
install:
	go mod tidy

## build: Build binary applications
build:
	templ generate
	npx tailwindcss -i ./static/main.css -o ./static/tailwind.css
	go build -o bin/goweb main.go

## run: Run goweb 
run: build
	./bin/goweb svc --config=$(config)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run with parameter options: "
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo