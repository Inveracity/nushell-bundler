# Nushell Bundler

> ⚠️ This is purely for fun

This is an example of how to create a binary out of nushell scripts using Go.

This means writing a whole command-line program in Nushell and then running it as an executable.

## How does it work?

Go can embed static files into the binary at build time.

`main.go` simply takes the `nu` binary and nushell scripts put in the `/src` directory.

At runtime the Go program will write the nushell scripts to `/tmp/gonu*` but will execute nu from memory with the flag `-n` meaning it uses default Nu configurations.

## Features

- Nested modules
- Pass commandline arguments

## How to

Write some code in `src/main.nu`

build it

```sh
./toolkit build mytool # outputs: mytool

# try commandline argument
./mytool --help
./mytool --blue
./mytool boop
```

## Requirements

| bin     | version     |
|---------|-------------|
| Go      | >= 1.20     |
| Nushell | any version |
