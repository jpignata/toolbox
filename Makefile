.PHONY: install clean

install: bin/aoc bin/gist
	cp bin/* ~/bin

clean:
	rm -f bin/*

bin/aoc:
	go build -o bin/aoc aoc/main.go

bin/gist:
	go build -o bin/gist gist/main.go
