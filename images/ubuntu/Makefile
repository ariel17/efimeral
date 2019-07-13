IMAGE = ariel17/efimeral:ubuntu-19.04
.build:
	docker build . -t $(IMAGE)

.push:
	docker push $(IMAGE)

.up:
	docker run -p 4200:4200 $(IMAGE)

default: .build .push
