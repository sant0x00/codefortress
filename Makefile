LOCAL_REPO = $(HOME)/Development/github.com/rios0rios0/pipelines

.PHONY: build
build:
	rm -rf bin
	go mod tidy
	go build -o bin/codefortress ./cmd/codefortress

.PHONY: run
run:
	cd bin; ./codefortress

.PHONY: pipelines
pipelines:
	@ cd scripts; \
 	chmod +x pipelines.sh; \
	./pipelines.sh

.PHONY: lint
lint:
	make pipelines
	${LOCAL_REPO}/global/scripts/golangci-lint/run.sh .

.PHONY: install
install:
	make build
	sudo cp -v bin/codefortress /usr/local/bin/codefortress
