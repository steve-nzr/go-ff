all: update deploy

.PHONY: update
update:
	@go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod

.PHONY: deploy
deploy:
	ibazel run //deployments/kubernetes:go-ff.apply
