all: refresh build

.PHONY: refresh
refresh: refresh-deps refresh-app

.PHONY: refresh-deps
refresh-deps:
	@go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod

.PHONY: refresh-app
refresh-app:
	@find . -name "./api/**/*.pb.*" -type f -delete
	bazel run //:gazelle

.PHONY: build
build: build-action build-chat build-cluster build-connectionserver build-entity build-login build-moving

.PHONY: build-action
build-action:
	bazel run //cmd/action:go_default_image -- --norun

.PHONY: build-chat
build-chat:
	bazel run //cmd/chat:go_default_image -- --norun

.PHONY: build-cluster
build-cluster:
	bazel run //cmd/cluster:go_default_image -- --norun

.PHONY: build-connectionserver
build-connectionserver:
	bazel run //cmd/connectionserver:go_default_image -- --norun

.PHONY: build-entity
build-entity:
	bazel run //cmd/entity:go_default_image -- --norun

.PHONY: build-login
build-login:
	bazel run //cmd/login:go_default_image -- --norun

.PHONY: build-moving
build-moving:
	bazel run //cmd/moving:go_default_image -- --norun

.PHONY: run-action
run-action:
	ibazel run //cmd/action:go_default_image

.PHONY: run-chat
run-chat:
	ibazel run //cmd/chat:go_default_image

.PHONY: run-cluster
run-cluster:
	ibazel run //cmd/cluster:go_default_image

.PHONY: run-connectionserver
run-connectionserver:
	ibazel run //cmd/connectionserver:go_default_image

.PHONY: run-entity
run-entity:
	ibazel run //cmd/entity:go_default_image

.PHONY: run-login
run-login:
	ibazel run //cmd/login:go_default_image

.PHONY: run-moving
build-moving:
	ibazel run //cmd/moving:go_default_image
