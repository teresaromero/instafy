.PHONY: build
build:
	go build -o bin/instafy main.go

.PHONY: run
run:
	go run main.go

.PHONY: broot
broot:
	bin/instafy