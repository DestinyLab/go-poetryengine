/*
Package poetryengine is the engine of lottery poetry.

Example:
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
*/
package poetryengine

import (
	"math/rand"
	"time"
)

// A Poem returns lottery-poetry.
type Poem struct {
	Title string   `json:"title"`
	Gua   []string `json:"gua"`
	Poem  []string `json:"poem"`
	Item  []string `json:"item"`
}

// A Suit returns slice of Poem.
type Suit struct {
	list []Poem
}

// New creates an instance of Suit.
func New(list []Poem) *Suit {
	return &Suit{list: list}
}

// Get finds poem by ID.
func (s *Suit) Get(poemid int) Poem {
	return s.list[poemid]
}

// Total returns total number of Suit.
func (s *Suit) Total() int {
	return len(s.list)
}

// GetList returns all poems.
func (s *Suit) GetList() []Poem {
	return s.list
}

// Draw picks random poem.
func (s *Suit) Draw() Poem {
	rand.Seed(time.Now().Unix())

	return s.list[rand.Intn(len(s.list))]
}
