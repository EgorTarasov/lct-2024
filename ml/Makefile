.PHONY:python-create-grpc
python-create-grpc:
	python -m grpc_tools.protoc \
	--python_out=./stubs \
	--pyi_out=./stubs \
	--grpc_python_out=./stubs \
	--proto_path ../proto inference.proto





#.PHONY: docker-build
#docker-build:
#	docker build \
#		--tag ghcr.io/egortarasov/lct-2024/ml:$(VERSION) \
#		--platform linux/amd64 .
#
#
#
#.PHONY: docker-push
#docker-push:
#	docker push ghcr.io/egortarasov/lct-2024/ml:$(VERSION)
#
#
#.PHONY: docker-publish
#docker-publish: docker-build docker-push

VERSION := 0.1


.PHONY: docker-build-amd64
docker-build-amd64:
	docker build \
		--tag ghcr.io/egortarasov/lct-2024/ml:amd64-$(VERSION) \
		--platform linux/amd64 .

.PHONY: docker-build-arm64
docker-build-arm64:
	docker build \
		--tag ghcr.io/egortarasov/lct-2024/ml:arm64-$(VERSION) \
		--platform linux/arm64 .



.PHONY: docker-push-amd64
docker-push-amd64:
	docker push ghcr.io/egortarasov/lct-2024/ml:amd64-$(VERSION)

.PHONY: docker-push-arm64
docker-push-arm64:
	docker push ghcr.io/egortarasov/lct-2024/ml:arm64-$(VERSION)

.PHONY: docker-publish-amd64
docker-publish-amd64: docker-build-amd64 docker-push-amd64

.PHONY: docker-publish-arm64
docker-publish-arm64: docker-build-arm64 docker-push-arm64

.PHONY: docker-publish
docker-publish:	docker-publish-amd64 docker-publish-arm64