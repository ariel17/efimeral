.PHONY: build
.DEFAULT_GOAL := default
IMAGE = ariel17/efimeral:web

install:
	yarn install

build: install
	yarn run build

deploy: install
	yarn run deploy

image-build:
	docker build . -t $(IMAGE)

image-push:
	docker push $(IMAGE)

up: install
	yarn start

clean: 
	rm -rf build

default: deploy image-build image-push
