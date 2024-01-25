# Go-File-Sync

Go-File-Sync is a simple command-line tool that allows you to synchronize files between a server and a client.

## Installation

To install Go-File-Sync, you need to have [Go](https://go.dev/) installed on your system. Then, you can use the following command to download and build the executable:

```bash
go get github.com/Jxpro/Go-File-Sync
```

This will create a Go-File-Sync binary in your `$GOPATH/bin` directory. You can also use `go install` to install the binary to a specific location.

## Usage

Go-File-Sync has two modes: server and client. You can specify the mode using the `-m` or `--mode` flag. For example:

```bash
Go-File-Sync -m server
```

or

```bash
Go-File-Sync --mode client
```

The server mode will start a file server that listens for incoming connections from clients. The server will store the files in a directory named `files` in the current working directory. You can change the port number by setting the `PORT` environment variable. The default port is `8080`.

The client mode will start a file client that connects to a server and synchronizes the files in a directory named `files` in the current working directory. You can change the server address by setting the `SERVER` environment variable. The default address is `localhost:8080`.

To synchronize the files, the client will send a list of file names and hashes to the server, and the server will respond with a list of files that need to be updated. The client will then send or receive the files accordingly.