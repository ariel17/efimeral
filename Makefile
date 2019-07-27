.DEFAULT_GOAL := default

image-up:
	docker-compose up -d && \
		docker-compose logs -f

image-clean:
	docker-compose rm -fsv

default:
	$(MAKE) -C api
	$(MAKE) -C web
	$(MAKE) -C images
	$(MAKE) -C docs
