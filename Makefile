.PHONY: api web images
api:
	cd api && $(MAKE) all

web:
	cd web && $(MAKE) all

images:
	cd images && $(MAKE) all

up:
	docker-compose up -d && \
		docker-compose logs -f

clean:
	docker-compose rm -fsv

all: api web images
