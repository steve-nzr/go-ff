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
build: build-action-image build-chat-image build-cluster-image build-connectionserver-image build-entity-image build-login-image build-moving-image

.PHONY: build-action-image
build-action-image:
	bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/action:image -- --norun

.PHONY: build-chat-image
build-chat-image:
	bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/chat:image -- --norun

.PHONY: build-cluster-image
build-cluster-image:
	bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/cluster:image -- --norun

.PHONY: build-connectionserver-image
build-connectionserver-image:
	bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/connectionserver:image -- --norun

.PHONY: build-entity-image
build-entity-image:
	bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/entity:image -- --norun

.PHONY: build-login-image
build-login-image:
	bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/login:image -- --norun

.PHONY: build-moving-image
build-moving-image:
	bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/moving:image -- --norun
