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

image: build
	docker build . -t ariel17/efimeral:api && \
		docker push ariel17/efimeral:api

up: build
	docker run -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock ariel17/efimeral:api

clean:
	rm api && rm docs/*.html

all: image
