all: refresh-deps refresh-app

.PHONY: refresh-deps
refresh-deps:
	@go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod

.PHONY: refresh-app
refresh-app:
	@find . -name "./api/**/*.pb.*" -type f -delete
	bazel run //:gazelle
