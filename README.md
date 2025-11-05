# httpshare

**httpshare** is a lightweight, web-based file explorer written in Go. It allows you to quickly share and browse files over HTTP with minimal setup.

[![Releases](https://img.shields.io/github/v/release/isa0-gh/httpshare)](https://github.com/isa0-gh/httpshare/releases)

## Features

- Lightweight and easy to use
- Simple web interface for browsing and downloading files
- Configurable port and directory
- Cross-platform: works on Windows, Linux, and macOS

## Installation

Make sure you have [Go](https://golang.org/dl/) installed. Then, run:

```bash
go install github.com/isa0-gh/httpshare@latest
````

This installs the `httpshare` binary in your Go `bin` directory (usually `$GOPATH/bin`).

## Usage

Start the file explorer with the following command:

```bash
httpshare [--port <port>] [--directory <path>]
```

### Options

* `--port`: Set the port to serve the files (default: `8080`)
* `--directory`: Specify the directory to serve (default: current directory)

Then, open your web browser and navigate to:

```
http://localhost:<port>
```

You will see a simple web interface to explore and download files in the specified directory.

### Example

Serve files from `/home/user/Documents` on port 3000:

```bash
httpshare --port 3000 --directory /home/user/Documents
```

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License.
