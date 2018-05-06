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
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -o $(PROJECT)

docker-build: $(PROJECT)
	docker build .

deploy-local: docker-build
	ENV=local-dev docker-compose -p webapp up -d

.PHONY: deploy-local
