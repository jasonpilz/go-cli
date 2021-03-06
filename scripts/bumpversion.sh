#!/usr/bin/env bash

set -euo pipefail

semver() {
  git tag -l | sort --version-sort | tail -n1 | cut -c 2-
}

ORIGIN=${ORIGIN:-origin}

# Bump defaults to patch, providing aliases for patch, minor and major
BUMP=${BUMP:-patch}
case "$BUMP" in
  feature | minor)
    BUMP="minor"
    ;;
  breaking | major)
    BUMP="major"
    ;;
  *)
    BUMP="patch"
    ;;
esac

if [[ $(git status --porcelain) != "" ]]; then
  echo "Error: repo is dirty. Run git status, clean repo and try again."

  exit 1
elif [[ $(git status --porcelain -b | grep -e "ahead" -e "behind") != "" ]]; then
  echo "Error: repo has unpushed commits. Push commits to remote and try again."

  exit 1
fi

version=$(semver)
new_version="v$(sembump --kind "$BUMP" "$version")"

echo "==> Bumping version from v${version} to ${new_version}"

git tag -m "release ${new_version}" -a "$new_version" && git push "${ORIGIN}" tag "$new_version"
