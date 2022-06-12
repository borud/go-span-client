all: vet lint test build

build: coap

coap:
	@cd cmd/$@ && go build -o ../../bin/$@ --trimpath -tags osusergo,netgo

test:
	@echo "*** $@"
	@go test ./...

vet:
	@echo "*** $@"
	@go vet ./...

lint:
	@echo "*** $@"
	@revive ./... 