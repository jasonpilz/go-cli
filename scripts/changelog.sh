#!/usr/bin/env bash

set -ou

# TEMPLATE_TODO: Update ORG and REPO variables
ORG=jasonpilz
REPO=go-cli

tfile=$(mktemp /tmp/$REPO-CHANGELOG-XXXXXX)
github-changelog-generator -org "$ORG" -repo "$REPO" >"$tfile"

echo "$tfile"
