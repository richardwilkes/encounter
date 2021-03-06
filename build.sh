#!/usr/bin/env bash
set -eo pipefail

TARGET_EXE="$GOPATH/bin/$(basename "$(pwd)")"
VERSION=0.1

# Setup GIT_VERSION
if command -v git 2>&1 > /dev/null; then
  if [ -z "$(git status --porcelain)" ]; then
    STATE=clean
  else
    STATE=dirty
  fi
  GIT_VERSION=$(git rev-parse HEAD)-$STATE
else
  GIT_VERSION=Unknown
fi

# Clean out old information
find . \( -iname "*_gen.go" -o -iname ".DS_Store" \) -exec rm \{\} \;

# Generate some code
go generate ./...

# Build the executable
LINK_FLAGS="-X github.com/richardwilkes/toolbox/cmdline.AppVersion=$VERSION"
LINK_FLAGS="$LINK_FLAGS -X github.com/richardwilkes/toolbox/cmdline.GitVersion=$GIT_VERSION"
go install -v -ldflags=all="$LINK_FLAGS"

# Add the internal filesystem to the executable
cd internal/assets
/bin/rm -f assets.zip
zip -q -D -r -9 assets.zip dynamic static
cd ../..
cat internal/assets/assets.zip >> "$TARGET_EXE"
/bin/rm -f internal/assets/assets.zip
zip -q -A "$TARGET_EXE"
