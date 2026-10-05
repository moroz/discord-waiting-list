#!/usr/bin/env -S bash -euo pipefail
# Builds rel/release.tar.gz: the server binary for GOOS/GOARCH, with its built
# CSS bundle (assets/dist, served straight from the binary's working directory
# by handlers/router.go) and the goose migrations. Follows medic-go's
# deploy/build-release.sh.

gitroot="$(git rev-parse --show-toplevel)"

export GOOS="${GOOS:-freebsd}"
export GOARCH="${GOARCH:-amd64}"
export CGO_ENABLED=0

cd "$gitroot"
rm -rf rel
mkdir -p rel/assets

pnpm install --frozen-lockfile
pnpm run build
go build -trimpath -ldflags="-s -w" -o rel/server .

cp -R assets/dist rel/assets/
cp -R db/migrations rel/

tar_opts="--no-xattrs"
if [[ "$(uname)" = "Darwin" ]]; then
  tar_opts="$tar_opts --no-mac-metadata"
fi
cd rel && tar czf release.tar.gz $tar_opts server assets migrations
