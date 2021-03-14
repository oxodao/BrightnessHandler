build:
	go mod vendor
	go build -o bh .
	sudo chown root:root bh
	sudo chmod u+s bh

install:
	mv bh /usr/bin
