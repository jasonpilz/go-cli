#!make
include .env.goreleaser
export

.PHONY: install
install:
	@echo "==> Installing cli locally"
	@echo "==> TEMPLATE_TODO: Rename 'cli' folder to the desired name of cli binary"
	go install ./cmd/cli

.PHONY: docker/login
docker/login:
	@echo "==> Logging in to docker registry"
	@echo "==> TODO: Configuration required for registry access"
	# $$(aws ecr get-login --no-include-email --region us-east-1)

.PHONY: lint/shellcheck
lint/shellcheck:
	@echo "==> Linting shell scripts"
	@scripts/shellcheck.sh

.PHONY: lint/vet
lint/vet:
	@echo "==> Linting with go vet"
	@go vet ./...

.PHONY: git/tag
git/tag: _os/install/sembump
	@echo "==> Bumping version BUMP=${BUMP} tag"
	@ORIGIN=${ORIGIN} scripts/bumpversion.sh

.PHONY: git/changes
git/changes: _os/install/gclg
	@echo "==> List merged PRs since last release"
	@changes=$(shell GITHUB_TOKEN=${GITHUB_TOKEN} scripts/changelog.sh) && cat $$changes && rm -f $$changes

.PHONY: github/changelog
github/changelog: _os/install/gclg
	@scripts/changelog.sh

.PHONY: github/release
github/release:
	@echo "==> Releasing"
	@scripts/release.sh

.PHONY: _os/install/gclg
_os/install/gclg:
	@GO111MODULE=off go get -u github.com/digitalocean/github-changelog-generator

.PHONY: _os/install/sembump
_os/install/sembump:
	@echo "==> Installing/updating sembump package"
	@GO111MODULE=off go get -u github.com/jessfraz/junk/sembump
