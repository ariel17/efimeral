.PHONY: build
install:
	yarn install

deploy: install
	yarn run deploy

image:
	docker build . -t ariel17/efimeral:web && \
		docker push ariel17/efimeral:web

up: install
	yarn start

clean: 
	rm -rf build

all: deploy image
