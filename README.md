# GO CLI

> Template for building command-line interfaces in golang.

## Template Use

This repo is intended to be used as [a template](https://github.blog/2019-06-06-generate-new-repositories-with-repository-templates/)
for building go-based command-line interfaces.

See the [github docs](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template)
For more info.

After creating a new project from this template, some customization will be required.
To aid in identifying and locating these items, comments including the words `TEMPLATE_TODO`
have been added to the codebase. Using your editor of choice, search/grep for this tag
and make adjustments as indicated by comments at those locations. These comments can
then be deleted if desired.

### Releasing

> This project follows [semantic versioning](semver.org).

When it's time to create a new release version:

1. `make git/changes` to review changes since the last release.
   Based on the changes, decide what kind of release is required (bugfix, feature or breaking).

2. `BUMP=(major|minor|patch|bugfix|feature|breaking) make git/tag` to create a new git tag.
   (bugfix, feature and breaking are aliases for semver's patch, minor and major).
   BUMP defaults to patch. The command assumes you have a remote repository named
   `origin` pointing to this repository.  If you'd prefer to specify a different remote
   repository, you can do so by setting `ORIGIN=(preferred remote name)`.

3. `make github/release` to create a new release based on the most recent git tag.
   This step will build application binaries for each os and architecture, docker images
   based on the binary, upload the images to a docker registry and the binaries to the github release.

### Prerequisites

* [goreleaser](https://goreleaser.com/install/)
* [docker](https://docs.docker.com/install/)
* `GITHUB_TOKEN` defined in `.env.goreleaser` with `repo` access.
