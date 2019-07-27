IMAGE = ariel17/efimeral:archlinux-latest
.build:
	docker build . -t $(IMAGE)

.push:
	docker push $(IMAGE)

.up:
	docker run -p 4200:4200 $(IMAGE)

default: .build .push
