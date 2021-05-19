[![GoDoc](https://godoc.org/github.com/mafulong/godal?status.svg)](https://pkg.go.dev/github.com/mafulong/godal)
[![Codecov.io](https://codecov.io/github/mafulong/godal/coverage.svg?branch=main)](https://codecov.io/github/mafulong/godal?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/mafulong/godal)](https://goreportcard.com/report/github.com/mafulong/godal)

[comment]: <> ([![Build Status]&#40;https://travis-ci.org/mafulong/godal.svg?branch=main&#41;]&#40;https://travis-ci.org/mafulong/godal&#41;)
godal
===

godal provides the ability to generate specific golang code. The goal is to enable developers to write fast code in an
expressive way.

Functions:

1. Generate orm model corresponding to golang by specifying sql ddl file, which can be used by gorm.
2. Update later

## Usage Documentation

## Installation

Using this package requires a working Go
environment. [See the install instructions for Go](http://golang.org/doc/install.html).

Go Modules are required when using this
package. [See the go blog guide on using Go Modules](https://blog.golang.org/using-go-modules).

How to install godal?

```shell
git clone git@github.com:mafulong/godal.git
cd godal
go install
```

Make sure that you can have result when call commend`which godal`

## SQL to golang model

```shell
godal gen model --database {your databaseName} {your sqlFile}
```

### Example

You can enter the test directory in this repo.

```shell
godal gen model --database testdb gen_model.sql
```

[What you can get?](https://github.com/mafulong/godal/test/model/)

example file1:

```go
package testdb

import "time"

type TestTb1 struct {
	TestId         int       `gorm:"Column:test_id" json:"test_id"`
	TestTitle      string    `gorm:"Column:test_title" json:"test_title"`
	TestAuthor     string    `gorm:"Column:test_author" json:"test_author"`
	SubmissionDate time.Time `gorm:"Column:submission_date" json:"submission_date"`
}

func (TestTb1) TableName() string {
	return "test_tb1"
}

```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.