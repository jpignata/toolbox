# Toolbox

Random command line tools I use day to day. Currently porting from various shell, Python, and Ruby scripts to Go for portability.

## Tools

### aoc

Fetches input for an [Advent of Code][1] puzzle and prints it to `STDOUT`. If year is omitted, it defaults to the current year.

#### Usage

`aoc day [year]`

```console
$ aoc 1 2016 | tee input.txt
R2, L1, R2, R1, R1, L3, R3, L5, L5, L2, L1, R4, R1, R3, L5, L5, R3, L4, L4, R5, R4, R3, L1, L2, R5, R4, L2, R1, R4, R4, L2, L1, L1, R190, R3, L4, R52, R5, R3, L5, R3, R2, R1, L5, L5, L4, R2, L3, R3, L1, L3, R5, L3, L4, R3, R77, R3, L2, R189, R4, R2, L2, R2, L1, R5, R4, R4, R2, L2, L2, L5, L1, R1, R2, L3, L4, L5, R1, L1, L2, L2, R2, L3, R3, L4, L1, L5, L4, L4, R3, R5, L2, R4, R5, R3, L2, L2, L4, L2, R2, L5, L4, R3, R1, L2, R2, R4, L1, L4, L4, L2, R2, L4, L1, L1, R4, L1, L3, L2, L2, L5, R5, R2, R5, L1, L5, R2, R4, R4, L2, R5, L5, R5, R5, L4, R2, R1, R1, R3, L3, L3, L4, L3, L2, L2, L2, R2, L1, L3, R2, R5, R5, L4, R3, L3, L4, R2, L5, R5
```

### gist

Posts the given files and/or input from `STDIN` as a [GitHub Gist][5]. Gists are private by default, but can be made public via `-p`. Returns a link to the Gist.

#### Usage

`gist [-f <filename>] [-d <description>] [-n <name of stdin file>] [-p]`

```console
$ cat coolfile.txt | gist
https://gist.github.com/jpignata/0123456789abdefc0123456789abcdef

$ ./some-program | gist -n output.txt
https://gist.github.com/jpignata/0123456789abdefc0123456789abcdef

$ gist -f ./file.txt -d "Here's a file I have" -p
https://gist.github.com/jpignata/0123456789abdefc0123456789abcdef
```

### bitly

Shortens the given link and returns a [Bitlink][6].

#### Usage

`bitly [url]`

```console
$ bitly https://www.audible.com/pd/The-Three-Body-Problem-Audiobook/B00P027
https://adbl.co/2WPs8b7
```

## Dependencies

Tools that require authentication use [AWS System Manager][2] [Parameter Store][3] to fetch credentials. See [pkg/ssm/secure_string.go][4] for details.

[1]: https://www.adventofcode.com
[2]: https://aws.amazon.com/systems-manager/
[3]: https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html
[4]: pkg/ssm/secure_string.go
[5]: https://gist.github.com
[6]: https://bit.ly
