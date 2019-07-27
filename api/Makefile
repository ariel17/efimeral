.DEFAULT_GOAL := default
IMAGE = ariel17/efimeral:api

api-install:
	dep ensure

api-build: api-install
	go test ./... && \
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docs-install:
	cd docs && yarn install

docs-build: docs-install
	cd docs && ./node_modules/.bin/aglio -i api.apib -o api.html

build: api-build docs-build

image-build: build
	docker build . -t $(IMAGE)

image-push:
	docker push $(IMAGE)

image-up: image-build
	docker run -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock $(IMAGE)

up:
	ENVIRONMENT=production BASE_URL=http://localhost GIN_MODE=release go run api.go

up-development:
	ENVIRONMENT=development BASE_URL=http://localhost go run api.go

clean:
	rm api && rm docs/*.html

default: image-build image-push
