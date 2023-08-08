# yc

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.20-%23007d9c)
[![GoDoc](https://godoc.org/github.com/denisdubochevalier/yc?status.svg)](https://pkg.go.dev/github.com/denisdubochevalier/yc)
![Build Status](https://github.com/denisdubochevalier/yc/actions/workflows/go.yml/badge.svg)
![Lint Status](https://github.com/denisdubochevalier/yc/actions/workflows/golangci-lint.yml/badge.svg)
![CodeQL Status](https://github.com/denisdubochevalier/yc/actions/workflows/codeql.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/denisdubochevalier/yc)](https://goreportcard.com/report/github.com/denisdubochevalier/yc)
[![Coverage](https://img.shields.io/codecov/c/github/denisdubochevalier/yc)](https://codecov.io/gh/denisdubochevalier/yc)
[![License](https://img.shields.io/github/license/denisdubochevalier/yc)](./LICENSE)

A short library to implement a Y combinator, defined as:

$$\lambda{f}.(\lambda{x}.f(xx))(\lambda{x}.f(xx))$$
