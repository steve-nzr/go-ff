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
	bazel run //deployments/kubernetes/action:action.apply

.PHONY: build-chat
build-chat:
	bazel run //deployments/kubernetes/chat:chat.apply

.PHONY: build-cluster
build-cluster:
	bazel run //deployments/kubernetes/cluster:cluster.apply

.PHONY: build-connectionserver
build-connectionserver:
	bazel run //deployments/kubernetes/connectionserver:connectionserver.apply

.PHONY: build-entity
build-entity:
	bazel run //deployments/kubernetes/entity:entity.apply

.PHONY: build-login
build-login:
	bazel run //deployments/kubernetes/login:login.apply

.PHONY: build-moving
build-moving:
	bazel run //deployments/kubernetes/moving:moving.apply

.PHONY: run-action
run-action:
	ibazel run //deployments/kubernetes/action:action.apply

.PHONY: run-chat
run-chat:
	ibazel run //deployments/kubernetes/chat:chat.apply

.PHONY: run-cluster
run-cluster:
	ibazel run //deployments/kubernetes/cluster:cluster.apply

.PHONY: run-connectionserver
run-connectionserver:
	ibazel run //deployments/kubernetes/connectionserver:connectionserver.apply

.PHONY: run-entity
run-entity:
	ibazel run //deployments/kubernetes/entity:entity.apply

.PHONY: run-login
run-login:
	ibazel run //deployments/kubernetes/login:login.apply

.PHONY: run-moving
run-moving:
	ibazel run //deployments/kubernetes/moving:moving.apply
