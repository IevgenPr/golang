# Learn Go

[![Build Status](https://travis-ci.org/IevgenPr/golang.svg?branch=master)](https://travis-ci.org/IevgenPr/golang)

Inspired by: [Gitbook](https://quii.gitbook.io/learn-go-with-tests)

Go install: brew install go
GO update: brew upgrade go

Create module:
/src/golang/src/github.com/ipr0/golang

# go mod init github.com/ipr0/golang

git repo is here:
~/src/golang/src/github.com/ipr0/golang

- A Go module must be a VCS repository or a VCS repository should contain a single Go module.
- A Go module should contain one or more packages
- A package should contain one or more .go files in a single directory.

Use modules, do not use workspace (latter is unable to control package versions)
https://blog.golang.org/using-go-modules
