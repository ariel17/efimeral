.PHONY: api web images
api:
	cd api && $(MAKE) all

web:
	cd web && $(MAKE) all

images:
	cd images && $(MAKE) all

all: api web images
