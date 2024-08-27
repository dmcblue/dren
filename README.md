# dren

## Go

This project mostly uses the [Go](https://go.dev/) language.

### Version Manager

It is suggested to use [g](https://github.com/voidint/g) as a Go version manager.

The current version of Go is contained in `.gorc`

You can install the current version with
```shell
g install $(cat .gorc)
```

If the version has been updated, you can update the `.gorc` file with (Bash):
```bash
GVER="$(go version)"; regex="([0-9]+\.[0-9]+\.[0-9]+)"; if [[ "$GVER" =~ $regex ]]; then echo "${BASH_REMATCH[1]}"; fi > .gorc
```

## Running

```shell
go run *.go
```
