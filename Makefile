build:
	go build -o server .

release:
	rm -rf rel/*
	pnpm run build
	go build -o rel/server
	cp -R assets rel/
	cp -R db/migrations rel/
	cd rel && tar czf release.tar.gz server assets/ migrations/
