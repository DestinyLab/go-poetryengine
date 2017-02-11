package poetryengine

import "fmt"
import "testing"

var s *Suit

var p = []Poem{
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

func init() {
	s = New(p)
}

func TestNew(t *testing.T) {
	New(p)
}

func TestGet(t *testing.T) {
	fmt.Println(s.Get(0).Title)
}

func TestTotal(t *testing.T) {
	fmt.Println(s.Total())
}

func TestGetList(t *testing.T) {
	fmt.Printf("%T\n", s.GetList())
}

func TestDraw(t *testing.T) {
	fmt.Println(s.Draw().Title)
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(p)
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Get(1)
	}
}

func BenchmarkTotal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Total()
	}
}

func BenchmarkGetList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.GetList()
	}
}

func BenchmarkDraw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Draw()
	}
}
