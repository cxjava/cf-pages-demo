.PHONY: dev
dev:
	wrangler pages dev ./pages

.PHONY: build
build:
	go run github.com/syumai/workers/cmd/workers-assets-gen@v0.18.0 -mode=go
	GOOS=js GOARCH=wasm go build -o ./build/app.wasm .

.PHONY: deploy
deploy:
	wrangler pages deploy ./pages
