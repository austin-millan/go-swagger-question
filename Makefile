generate-base-extra:
	swagger generate server \
		--with-flatten=full \
		--name=base-server \
		--with-context \
		--spec=pkg/spec/base.yml \
		--server-package=pkg/server/base \
		--model-package=pkg/server/base/models \
		--api-package=operations
	swagger mixin --format=yaml pkg/spec/base.yml pkg/spec/extra.yml --output=pkg/full.yml --quiet --ignore-conflicts
	swagger generate server \
		--with-flatten=full \
		--name=base-extra-server \
		--with-context \
		--spec=pkg/full.yml \
		--server-package=pkg/server/base_extra \
		--model-package=pkg/server/base_extra/models \
		--api-package=operations
	goimports -e -w .
	go fmt ./...
	go mod tidy