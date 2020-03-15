.PHONY: install clean

all: bin/aoc bin/gist bin/bitly

install: all
	cp bin/* ~/bin

clean:
	rm -f bin/*

bin/aoc:
	go build -o bin/aoc aoc/main.go

bin/gist:
	go build -o bin/gist gist/main.go

bin/bitly:
	go build -o bin/bitly bitly/main.go
