VERSION := 0.0.1


.PHONY: docker-build
docker-build:
	docker build \
		--tag ghcr.io/egortarasov/lct-2024/ui:$(VERSION) \
		--platform linux/amd64 .



.PHONY: docker-push
docker-push:
	docker push ghcr.io/egortarasov/lct-2024/ui:$(VERSION)


.PHONY: docker-publish
docker-publish: docker-build docker-push