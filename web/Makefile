.PHONY: build
.DEFAULT_GOAL := default
IMAGE = ariel17/efimeral:web

install:
	yarn install

build: install
	yarn run build

deploy: install
	yarn run deploy

image-build-development:
	docker build . -t $(IMAGE)

image-build-production:
	docker build . --build-arg REACT_APP_API_URL=http://api.efimeral.com.ar \
		--build-arg REACT_APP_DOCS_URL=/docs/index.html \
		--build-arg PUBLIC_URL= -t $(IMAGE)

image-push: image-build-production
	docker push $(IMAGE)

up: install
	yarn start

clean: 
	rm -rf build node_modules

default: deploy image-push
