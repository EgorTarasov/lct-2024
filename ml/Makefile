.PHONY:python-create-grpc
python-create-grpc:
	python -m grpc_tools.protoc \
	--python_out=./stubs \
	--pyi_out=./stubs \
	--grpc_python_out=./stubs \
	--proto_path ../proto inference.proto


VERSION := 0.0.1


.PHONY: docker-build
docker-build:
	docker build \
		--tag ghcr.io/egortarasov/lct-2024/ml:$(VERSION) \
		--platform linux/amd64 .



.PHONY: docker-push
docker-push:
	docker push ghcr.io/egortarasov/lct-2024/ml:$(VERSION)


.PHONY: docker-publish
docker-publish: docker-build docker-push