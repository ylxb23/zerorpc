.PHONY: build clean fmt vet test bench cover runRpc runWeb benchInfoRandom benchInfoFix benchLogin

ARGS=
COVERPROFILE=coverage.txt
DEBUG=

build: vendor clean fmt vet test
	go build ./...

clean:
	rm -rf bin/
	if [ -f "coverage.txt" ] ; then rm coverage.txt; fi

fmt:
	go fmt ./...

vendor:
	go mod vendor

vet:
	go vet ./...

test:
	go test ./... -coverprofile=$(COVERPROFILE) -covermode atomic -v $(DEBUG) $(ARGS)

cover:
	$(eval COVERPREFILE += -coverprofile=coverage.out)
	go test ./... -cover $(COVERPREFILE) -race $(ARGS) $(DEBUG)
	go tool cover -html=coverage.out
	rm -f coverage.out

bench:
	go test -v ./ -test.bench CallRpc -test.count 1 -benchtime=60s

