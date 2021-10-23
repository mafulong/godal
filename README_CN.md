[![GoDoc](https://godoc.org/github.com/mafulong/godal?status.svg)](https://pkg.go.dev/github.com/mafulong/godal)
[![Codecov.io](https://codecov.io/github/mafulong/godal/coverage.svg?branch=main)](https://codecov.io/github/mafulong/godal?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/mafulong/godal)](https://goreportcard.com/report/github.com/mafulong/godal)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

godal
===

[English](README.md) | [中文](README_CN.md) 

godal 提供了生成特定 Golang 代码的能力， 旨在提高研发编码的效率。

功能：

1. 基于 Mysql 的建表语句快速生成对应 Golang 的 Model，可直接被 ORM 框架 GORM 使用。
2. 待更新，欢迎大家提需求，会尽快适配。

godal 已经被加入到 [awesome-go](https://awesome-go.com/), [MR link](https://github.com/avelino/awesome-go/pull/3857/)
## 使用教程

## 安装 godal

首先需要golang环境 [See the install instructions for Go](http://golang.org/doc/install.html). [See the go blog guide on using Go Modules](https://blog.golang.org/using-go-modules).

git clone 后 进行 go install

```shell
git clone git@github.com:mafulong/godal.git
cd godal
go install
```

可以使用命令`which godal` 检验是否已安装

## SQL 生成 Model

```shell
godal gen model --database {your databaseName} {your sqlFile}
```

### 例子

在本 repo 的 test 文件夹中有个测试 sql, 可以进行用来测试

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
