knmr
===

knmr - Command Line Utility written in Go

# DESCRIPTION

knmr is command-line interface tool in Go. 

# INSTALLATION

## Build from source

```bash
$ go get github.com/ken2mer/knmr
$ go install github.com/ken2mer/knmr
```

## EXAMPLES

```
knmr agent
```
* every 3s tick until signal receive

The `knmr agent` command has an '-d' option for debug output.

```
knmr -d agent
```
* log output in DEBUG level

# CONTRIBUTION

1. Fork ([https://github.com/ken2mer/knmr/fork](https://github.com/ken2mer/knmr/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request
