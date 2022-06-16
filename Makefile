BINARY_NAME=scenario
ENTRYPOINT_FILE=cmd/api/main.go

.PHONY: all
all: build

.PHONY: help
help:
	@echo "Make Routines:"
	@echo " - build                       	build everything"
	@echo " - run                         	runs already build excutable"
	@echo " - clean               		  		clean source directory"
	@echo " - test                        	test everything"
	@echo " - test_coverage               	test and exports coverage report"
	@echo " - vet                         	examines Go source code and reports suspicious constructs"
	@echo " - tidy                        	run go mod tidy"
	@echo " - fmt               		  			run go fmt"
	@echo " - lint               		  			run golangci-lint linter"
	@echo " - install_tools               	install dev tools"
	@echo " - gen_swagger            				generate the swagger spec from code comments"
	@echo " - serve_swagger            	  	opens the swagger ui with the generated specs"
	@echo " - swagger 											generates swagger spec and open swagger ui"


.PHONY: build
build:
	go build -o ${BINARY_NAME} ${ENTRYPOINT_FILE}

.PHONY: run
run:
	go run ${ENTRYPOINT_FILE}

.PHONY: clean
clean:
	go clean
	rm ${BINARY_NAME}
	rm coverage.out

.PHONY: test
test:
	go test ./...

.PHONY: test_coverage
test_coverage:
	go test ./... -coverprofile=coverage.out

.PHONY: vet
vet:
	go vet ${ENTRYPOINT_FILE}

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	gofumpt -l -w .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: install_tools
install_tools:
	go mod download
	go install github.com/go-swagger/go-swagger/cmd/swagger
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	go install mvdan.cc/gofumpt

.PHONY: generate_swagger
generate_swagger:
	swagger generate spec --scan-models -o ./docs/swagger.json

.PHONY: serve_swagger
serve_swagger:
	swagger serve -F swagger ./docs/swagger.json


.PHONY: generate_and_server_swagger
generate_and_server_swagger: generate_swagger serve_swagger

