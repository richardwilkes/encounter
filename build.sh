#!/usr/bin/env bash
set -eo pipefail

APP_NAME=encounter
APP_BUNDLE_NAME=Encounter
APP_VERSION=0.1.0
APP_VERSION_SHORT=0.1
BUNDLE_ID=com.trollworks.encounter
COPYRIGHT_YEARS=2018-2019
COPYRIGHT_OWNER="Richard A. Wilkes"

# Process args
for arg in "$@"
do
	case "$arg" in
		--dev|-d)  DEV_MODE=1 ;;
		--help|-h)
			echo "$0 [options]"
			echo "  -d, --dev   Development mode (live FS)"
			echo "  -h, --help  This help text"
			exit 0
			;;
		*) echo "Invalid argument: $arg"; BAIL=1 ;;
	esac
done
if [ ! -z $BAIL ]; then
	exit 1
fi

# Setup OS_TYPE
case $(uname -s) in
	Darwin*)  OS_TYPE=darwin ;;
	#Linux*)   OS_TYPE=linux ;;
	MINGW64*) OS_TYPE=windows ;;
	*)        echo "Unsupported OS"; false ;;
esac

# Setup GIT_VERSION
if which git 2>&1 > /dev/null; then
	if [ -z "$(git status --porcelain)" ]; then
		STATE=clean
	else
		STATE=dirty
	fi
	GIT_VERSION=$(git rev-parse HEAD)-$STATE
else
	GIT_VERSION=Unknown
fi

# Ensure the build number is set to something
if [ -z $BUILD_NUMBER ]; then
	BUILD_NUMBER=Unknown
fi

# Setup tools we need
TOOLS_DIR="$PWD/tools"
/bin/rm -rf "$TOOLS_DIR"
mkdir -p "$TOOLS_DIR"
go build -o "$TOOLS_DIR/cef" github.com/richardwilkes/cef
export PATH="$TOOLS_DIR:$PATH"
cef install

# Prepare platform-specific distribution bundle
cef dist \
	--bundle "$APP_BUNDLE_NAME" \
	--executable "$APP_NAME" \
	--icon AppIcon.icns \
	--release "$APP_VERSION" \
	--short-release "$APP_VERSION_SHORT" \
	--year "$COPYRIGHT_YEARS" \
	--owner "$COPYRIGHT_OWNER" \
	--id $BUNDLE_ID
case $OS_TYPE in
	darwin)
		touch "dist/macos/$APP_BUNDLE_NAME.app" # Causes Finder to refresh its state
		TARGET_EXE="dist/macos/$APP_BUNDLE_NAME.app/Contents/MacOS/$APP_NAME"
		;;
	windows)
		TARGET_EXE="dist/windows/$APP_NAME.exe"
		;;
	*)
		echo "Unsupported OS"
		false
		;;
esac

find . -iname "*_gen.go" -exec rm \{\} \;
find . -iname ".DS_Store" -exec rm \{\} \;
go generate -tags gen ./...
go build -o "$TARGET_EXE" -v -ldflags=all="-X github.com/richardwilkes/toolbox/cmdline.AppVersion=$APP_VERSION_SHORT -X github.com/richardwilkes/toolbox/cmdline.GitVersion=$GIT_VERSION -X github.com/richardwilkes/toolbox/cmdline.BuildNumber=$BUILD_NUMBER" ./main.go

if [ -z $DEV_MODE ]; then
	cd internal/assets
	/bin/rm -f assets.zip
	zip -q -D -r -9 assets.zip dynamic static
	cd ../..
	cat internal/assets/assets.zip >> "$TARGET_EXE"
fi