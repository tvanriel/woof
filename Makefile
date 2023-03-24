.PHONY: clean
clean:
	rm woof-linux-amd64 ||:
	rm woof-windows-amd64.exe ||:
	rm woof-linux-arm ||:

woof-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o woof-linux-amd64 main.go
woof-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o woof-windows-amd64.exe main.go
woof-linux-arm:
	GOOS=linux GOARCH=arm go build -o woof-linux-arm main.go

.PHONY: install
install: woof-linux-amd64
	mkdir -p ~/.local/bin
	mv woof-linux-amd64 ~/.local/bin/woof


.PHONY: all
all: clean woof-linux-amd64 woof-linux-arm woof-windows-amd64
