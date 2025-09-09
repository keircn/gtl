# gtl - Go Title Linter

Transforms text into properly capitalized titles according to the Chicago Manual of Style.

## Installation

### Releases

The latest stable release is always available [here](https://github.com/keircn/gtl/releases/latest)

### Build From Source

```sh
git clone https://github.com/keircn/gtl; cd gtl
go build -o gtl cmd/gtl/main.go
./gtl --version
# optional
cp ./gtl ~/.local/bin # or wherever you put your binaries
```

## Usage

```
Usage:
  gtl [options] [text]
  echo "text" | gtl [options]

Options:
  -h, --help     Show this help message
  -v, --version  Show version information

Examples:
  gtl "all your base are belong to us" # All Your Base Are Belong To Us
  echo "the quick brown fox" | gtl
```

## License

This project is subject to the terms of the [MIT License](./LICENSE).
