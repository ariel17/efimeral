.PHONY: build
install:
	yarn install

build: install
	yarn run build

deploy: install
	yarn run deploy

image: build
	docker build . -t ariel17/efimeral:web && \
		docker push ariel17/efimeral:web

up: install
	yarn start

clean: 
	rm -rf build

all: deploy image
