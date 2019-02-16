.PHONY: install clean

install: bin/aoc
	cp bin/* ~/bin

clean:
	rm bin/*

bin/aoc:
	go build -o bin/aoc aoc/main.go
