.PHONY: build
build:
	@ cd cmd/codefortress; go build -o ../../bin/codefortress

.PHONY: run
run:
	@ cd bin; ./codefortress
