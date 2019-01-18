.PHONY: build

build:
	go build -o build/url-encode main.go

clean:
	rm -rf build/
