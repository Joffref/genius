test:
	go test -v

build: test
	go build -o build/genius

run: build
	./build/genius
