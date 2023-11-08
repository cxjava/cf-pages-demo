.PHONY: dev
dev:
	wrangler pages dev ./pages

.PHONY: build
build:
	go run github.com/syumai/workers/cmd/workers-assets-gen@v0.18.0
	tinygo build -o ./build/app.wasm -target wasm -no-debug ./...
	# GOARCH=wasm GOOS=js go build -o ./build/app.wasm .

.PHONY: deploy
deploy:
	wrangler pages deploy ./pages