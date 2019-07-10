.ubuntu-build:
	docker build . -f Dockerfile.ubuntu -t ariel17/efimeral:ubuntu-19.04

.ubuntu-push:
	docker push ariel17/efimeral:ubuntu-19.04

.archlinux-build:
	docker build . -f Dockerfile.archlinux -t ariel17/efimeral:archlinux-latest

.archlinux-push:
	docker push ariel17/efimeral:archlinux-latest

.up:
	docker run -p 4200:4200 ariel17/efimeral:$$IMAGE

.build: .ubuntu-build .archlinux-build

.push: .ubuntu-push .archlinux-push

all: .build .push
