.PHONY: docs
api:
	go test ./... && \
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docs:
	cd docs && ./node_modules/.bin/aglio -i api.apib -o api.html

image: api
	docker build . -t ariel17/efimeral:api && \
		docker push ariel17/efimeral:api

run:
	docker run -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock ariel17/efimeral:api

clean:
	rm api && rm docs/*.html

all: docs image
