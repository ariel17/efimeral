ubuntu-19.04:
	docker build . -f Dockerfile.ubuntu -t ariel17/efimeral:ubuntu-19.04 && \
		docker push ariel17/efimeral:ubuntu-19.04

ubuntu: ubuntu-19.04

up:
	docker run -p 4200:4200 ariel17/efimeral:$$IMAGE

all: ubuntu
