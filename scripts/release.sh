#!/usr/bin/env bash

set -eu

if [[ ! -x $(command -v goreleaser) ]] ; then
  echo "==> Install https://goreleaser.com/install/ then run again"

  exit 1
fi

release_notes="$(make github/changelog)"

goreleaser --rm-dist --release-notes="${release_notes}"

rm -f "$release_notes"
