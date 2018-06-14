GIT_SHA=`git rev-parse --short HEAD || echo`
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=fibome
all: build test
build: deps
				$(GOBUILD) -o bin/$(BINARY_NAME) -v
build-linux: deps
				CGO_ENABLED=0 GOOS=linux GOOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME) -v
test:
				$(GOTEST) -v ./...
clean:
				$(GOCLEAN)
				rm -rf bin/$(BINARY_NAME)
run: build
				./bin/$(BINARY_NAME)
deps:
				$(GOGET) github.com/julienschmidt/httprouter
docker-build:
				docker build -t dyk0/fibome:$(GIT_SHA) -f Dockerfile $(CURDIR)
docker-push:
				docker tag dyk0/fibome:$(GIT_SHA) dyk0/fibome:latest
				docker push dyk0/fibome:$(GIT_SHA)
				docker push dyk0/fibome:latest
