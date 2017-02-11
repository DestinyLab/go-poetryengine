# go-poetryengine

poetryengine is the engine of lottery poetry.

[![GoDoc](https://godoc.org/github.com/DestinyLab/go-poetryengine?status.svg)](https://godoc.org/github.com/DestinyLab/go-poetryengine) [![Go Report Card](https://goreportcard.com/badge/github.com/DestinyLab/go-poetryengine)](https://goreportcard.com/report/github.com/DestinyLab/go-poetryengine) [![Build Status](https://travis-ci.org/DestinyLab/go-poetryengine.svg?branch=master)](https://travis-ci.org/DestinyLab/go-poetryengine) [![Coverage Status](https://coveralls.io/repos/github/DestinyLab/go-poetryengine/badge.svg?branch=master)](https://coveralls.io/github/DestinyLab/go-poetryengine?branch=master)

## Installation

```
go get -u github.com/DestinyLab/go-poetryengine
```

## Usage

```go
package main

import (
  "fmt"

  p "github.com/DestinyLab/go-poetryengine"
)

func main() {
  poem := []p.Poem{
    {
      Title: "第一籤",
      Gua:   []string{"上上", "○○○○○"},
      Poem:  []string{"彩鳳呈祥瑞", "麒麟降帝都", "禍除迎福至", "喜氣自然生"},
      Item:  []string{"禍消福至", "求官得位", "求財大利", "婚姻成就", "出行大吉", "占病得安", "作事大吉", "考試得意"},
    },
    {
      Title: "第二籤",
      Gua:   []string{"上平", "○●●●●"},
      Poem:  []string{"從革宜更變", "時來合動遷", "龍門魚躍出", "凡骨作神仙"},
      Item:  []string{"行事得利", "作事可成", "占訟和吉", "求官得位", "走失見近", "婚姻成吉", "運途漸吉", "作事如意"},
    },
  }

  s := p.New(poem)

  fmt.Printf("get one:\n%v\n\n", s.Get(0).Title)
  fmt.Printf("get all:\n%v\n\n", s.GetList())
  fmt.Printf("total: %v\n", s.Total())
  fmt.Printf("random one:\n%v\n", s.Draw().Title)
}
```
