build:
	go mod vendor
	go build -o bh .

install:
	mv bh /usr/bin
