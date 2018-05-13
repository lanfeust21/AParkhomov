PROJECT=webapp-app
IMAGE=webapp-app
HUB=apharkomov
all: deploy-local

SOURCE := $(shell find . -name '*.go')

GOOS=linux
GOARCH=amd64

clean:
	[ ! -f ./$(PROJECT) ] || rm ./$(PROJECT)

$(PROJECT): clean
	echo Building for $(GOOS)/$(GOARCH)
	#CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -o $(PROJECT)
    xgo   -targets=linux/amd64 -out $(PROJECT) .
docker-build: $(PROJECT)
	docker build .

deploy-local: docker-build
	ENV=dev docker-compose -p webapp up -d

.PHONY: deploy-local
