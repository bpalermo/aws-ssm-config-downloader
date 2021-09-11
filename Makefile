.PHONY: update-repos
update-repos: gazelle
	@bazelisk run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

.PHONY: gazelle
gazelle:
	@bazelisk run //:gazelle

.PHONY: build
build: update-repos
	@bazelisk build //...

.PHONY: test
test:
	@bazelisk test //...

.PHONY: clean
clean:
	@bazelisk clean

.PHONY: build-image
build-image:
	@bazelisk run //cmd/config-downloader:go_image

.PHONY: all
all: gazelle build test