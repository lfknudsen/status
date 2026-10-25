# Status

CLI utility for quickly displaying HTTP status code information.

## Dependencies
The Makefile install script uses my [install](https://gitlab.com/lfknudsen/install) programme,
but compiling it manually does not actually require this.

## Setup

```
go build
```

Move the resulting binary and .status-codes.csv to somewhere on your PATH.

## Usage

```
status <code>           Display the code, title, and description of the HTTP status
status <title>          with the given code or title.

status all              List all the HTTP status codes along with title and description.
```