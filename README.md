# Brazanation Documents - Golang

[![Build Status](https://img.shields.io/travis/brazanation/go-documents/master.svg?style=flat-square)](https://travis-ci.org/brazanation/go-documents)
[![Codecov branch](https://img.shields.io/codecov/c/github/brazanation/go-documents/master.svg?style=flat-square)](https://codecov.io/gh/brazanation/go-documents)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/brazanation/go-documents)
[![Go Report Card](https://goreportcard.com/badge/github.com/brazanation/go-documents?style=flat-square)](https://goreportcard.com/report/github.com/brazanation/go-documents)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://github.com/brazanation/go-documents/blob/master/LICENSE)

A Golang library to provide Brazilian Documents safer, easier and fun!

## Installation

go-documents requires Go 1.11 or later.

```
go get github.com/brazanation/go-documents
```

If you want to get an specific version, please use the example below:

```
go get gopkg.in/brazanation/go-documents.v0
```

## Development

### Requirements

- Install [Go](https://golang.org)

### Makefile
```sh
// Clean up
$ make clean

// Run tests and generates html coverage file
make cover

// Download project dependencies
make depend

// Format all go files
make fmt

// Run linters
make lint

// Run tests
make test
```

## Documentation

Read the full documentation at [https://godoc.org/github.com/brazanation/go-documents](https://godoc.org/github.com/brazanation/go-documents).

### CNPJ (cadastro nacional da pessoa jurídica)

Company Identification or National Register of Legal Entities

```go
doc, err := brdocs.NewCnpj("99999090910270")
if err != nil {
    panic(err)
}
println(doc.String()) // prints 99999090910270
println(doc.Format()) // prints 99.999.090/9102-70
```

### CNH (carteira nacional de habilitação)

National Driving License
```go
doc, err := brdocs.NewCnh("83592802666")
if err != nil {
    panic(err)
}
println(doc.String()) // prints 83592802666
println(doc.Format()) // prints 83592802666
```

### PIS/PASEP (programa de integração social e programa de formação do patrimônio do servidor público)

Social Integration Program and Training Program of the Heritage of Public Servant

```php
doc, err := brdocs.NewPisPasep("51823129491")
if err != nil {
    panic(err)
}
println(doc.String()) // prints 51823129491
println(doc.Format()) // prints 518.23129.49-1
```
## License

This project is released under the MIT licence. See [LICENSE](https://github.com/brazanation/go-documents/blob/master/LICENSE) for more details.
