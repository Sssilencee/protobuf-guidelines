SERVICE_NAME_PATH=service_name
SERVICE_NAME_VERSION=v1
ANOTHER_SERVICE_NAME_PATH=another_service_name
ANOTHER_SERVICE_NAME_VERSION=v1

.PHONY: service_name
service_name:
	docker run --rm -v $(shell pwd):/work -w /work bufbuild/buf generate \
	    ${SERVICE_NAME_PATH}/${SERVICE_NAME_VERSION}/service_name.proto \
		--output ${SERVICE_NAME_PATH}

.PHONY: another_service_name
another_service_name:
	docker run --rm -v $(shell pwd):/work -w /work bufbuild/buf generate \
	    ${ANOTHER_SERVICE_NAME_PATH}/${ANOTHER_SERVICE_NAME_VERSION}/another_service_name.proto \
		--output ${ANOTHER_SERVICE_NAME_PATH}

.PHONY: lint
lint:
	docker run --rm -v $(shell pwd):/work -w /work bufbuild/buf lint
