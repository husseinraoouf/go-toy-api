BINARY_NAME=scenario
ENTRYPOINT_FILE=cmd/api/main.go
MIGRATION_CMD_FILE=cmd/migration/main.go

.PHONY: all
all: build

.PHONY: help
help:
	@echo "Make Routines:"
	@echo " - build                       	build everything"
	@echo " - run                         	runs already build excutable"
	@echo " - database_migrate              build everything"
	@echo " - clean               		  	clean source directory"
	@echo " - test                        	test everything"
	@echo " - test_coverage               	test and exports coverage report"
	@echo " - vet                         	examines Go source code and reports suspicious constructs"
	@echo " - tidy                        	run go mod tidy"
	@echo " - fmt               		  	run go fmt"
	@echo " - lint                 			run golangci-lint linter"
	@echo " - install_tools               	install dev tools"
	@echo " - gen_swagger            		generate the swagger spec from code comments"
	@echo " - serve_swagger            	  	opens the swagger ui with the generated specs"
	@echo " - swagger 						generates swagger spec and open swagger ui"


.PHONY: build
build:
	go build -o ${BINARY_NAME} ${ENTRYPOINT_FILE}

.PHONY: run
run:
	go run ${ENTRYPOINT_FILE}

.PHONY: database_migrate
database_migrate:
	go run ${MIGRATION_CMD_FILE}

.PHONY: clean
clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -f coverage.out

.PHONY: test
test:
	go test ./...

.PHONY: test_coverage
test_coverage:
	go test ./... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out

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

.PHONY: gen_swagger
gen_swagger:
	swagger generate spec --scan-models -o ./docs/swagger.json

.PHONY: serve_swagger
serve_swagger:
	swagger serve -F swagger ./docs/swagger.json


.PHONY: swagger
swagger: gen_swagger serve_swagger

