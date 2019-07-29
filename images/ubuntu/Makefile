.DEFAULT_GOAL := default
IMAGE = ariel17/efimeral:ubuntu-19.04
build:
	docker build . -t $(IMAGE)

push:
	docker push $(IMAGE)

up:
	docker run -p 10443:10443 $(IMAGE)

default: build push
