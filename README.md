# Telesend [![Build Status](https://travis-ci.org/cikupin/telesend.svg?branch=master)](https://travis-ci.org/cikupin/telesend) [![GoDoc](https://godoc.org/github.com/cikupin/telesend?status.svg)](https://godoc.org/github.com/cikupin/telesend) [![Go Report Card](https://goreportcard.com/badge/github.com/cikupin/telesend)](https://goreportcard.com/report/github.com/cikupin/telesend) [![Maintainability](https://api.codeclimate.com/v1/badges/c7f74987c5e58170b1b4/maintainability)](https://codeclimate.com/github/cikupin/telesend/maintainability)

This is a simple implementation to send telegram message to private group or channel.

Documentation can be found at [godoc](https://godoc.org/github.com/cikupin/telesend).

#### Example

```go
package main

import (
	"github.com/cikupin/telesend"
)

func main() {
	url := "https://web.telegram.org/#/im?p=g123456"
	token := "3455345:dfngue4ht3478hunbdf"

	id, err := telesend.GetGroupID(url)
	if err != nil {
		panic(err)
	}

	bot := telesend.NewBot(token, id)
	bot.SendMessage("this is a testing message")
}
```
