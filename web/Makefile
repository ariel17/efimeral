.PHONY: build
build:
	npm run build

image: build
	docker build . -t ariel17/efimeral:web && \
		docker push ariel17/efimeral:web

all: image
