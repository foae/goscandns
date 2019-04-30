all: lint test install_goscandns

install_goscandns:
	go install github.com/foae/goscandns/cmd/goscandns

lint:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;
	which gometalinter; if [ $$? -ne 0 ]; then go get -u github.com/alecthomas/gometalinter && gometalinter --install; fi
	gometalinter --vendor --exclude=repos --disable-all --enable=golint ./...
	go vet ./...

test:
	go test -v -short -cover ./...

run: install_goscandns
	ENV="dev" \
	$(GOPATH)/bin/goscandns
