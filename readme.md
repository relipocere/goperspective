[![Go Reference](https://pkg.go.dev/badge/github.com/exsocial/goperspective.svg)](https://pkg.go.dev/github.com/exsocial/goperspective) [![Go Report Card](https://goreportcard.com/badge/github.com/exsocial/goperspective)](https://goreportcard.com/report/github.com/exsocial/goperspective) [![Code Coverage](https://img.shields.io/badge/coverage-79%25-green)]
# GO Perspective API Client
[Perspective API](https://developers.perspectiveapi.com/s/) is a free tool that allows you to analyze the sentiment of comments for such attributes as toxicity, threats, etc.
This library is an implementation of the API client in Go. 
[Attributes and supported languages](https://developers.perspectiveapi.com/s/about-the-api-attributes-and-languages).
Methods and fields can be found [here](https://developers.perspectiveapi.com/s/about-the-api-methods).

## Getting started
### Installing

```sh
go get github.com/exsocial/goperspective
```
### Usage

Import the package into your project.

```go
import "github.com/exsocial/goperspective"
```

Construct a new API client.

```go
client := gp.NewClient(os.Getenv("TOKEN"))
```

See Documentation and Examples below for more detailed information.