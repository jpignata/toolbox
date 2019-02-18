.PHONY: install clean

install: bin/aoc bin/gist bin/pf
	cp bin/* ~/bin

clean:
	rm -f bin/*

bin/aoc:
	go build -o bin/aoc aoc/main.go

bin/gist:
	go build -o bin/gist gist/main.go

bin/pf:
	go build -o bin/pf pf/main.go
