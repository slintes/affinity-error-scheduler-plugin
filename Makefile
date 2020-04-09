COMMONENVVAR=GOOS=linux GOARCH=amd64
BUILDENVVAR=CGO_ENABLED=0

.PHONY: all
all: build

.PHONY: build
build: gofmt govet
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o bin/kube-scheduler main.go

.PHONY: gofmt
gofmt:
	@echo "Running gofmt"
	gofmt -s -w `find . -path ./vendor -prune -o -type f -name '*.go' -print`

.PHONY: govet
govet:
	@echo "Running go vet"
	go vet ./...

.PHONY: image-build
image-build: build
	@echo "building image"
	docker build -f images/Dockerfile -t quay.io/slintes/ta-scheduler:latest .

.PHONY: image-push
image-push: image-build
	@echo "pushing image"
	docker push quay.io/slintes/ta-scheduler:latest
