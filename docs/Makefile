.DEFAULT_GOAL := default
IMAGE = ariel17/efimeral:docs

docs-install:
	yarn install

docs-build: docs-install
	./node_modules/.bin/aglio -i api.apib -o index.html

image-build:
	docker build . -t $(IMAGE)

image-push: image-build
	docker push $(IMAGE)

up: image-build
	docker run -p 9090:80 $(IMAGE)

clean:
	rm -rf node_modules *.html

default: docs-build image-build image-push
