#!/usr/bin/env -S bash -euo pipefail
# Builds rel/release.tar.gz: the server binary for GOOS/GOARCH (the templates and
# countries.json are embedded in it), the Tailwind bundle in assets/dist, which the
# server serves itself under /assets/, and the migrations, which goose runs on the
# server before each start. Follows holy-shit-api's deploy/build-release.sh.

gitroot="$(git rev-parse --show-toplevel)"

export GOOS="${GOOS:-freebsd}"
export GOARCH="${GOARCH:-amd64}"
# modernc.org/sqlite is pure Go
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
