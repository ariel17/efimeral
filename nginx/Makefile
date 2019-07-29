.DEFAULT_GOAL := default
IMAGE = ariel17/efimeral:nginx

image-build:
	docker build . -t $(IMAGE)

image-push: image-build
	docker push $(IMAGE)

up: image-build
	docker run -p 80:80 $(IMAGE)

default: image-push
