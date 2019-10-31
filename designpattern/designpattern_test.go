package designpattern

import (
	"github.com/ryo-chin/algo-labo/designpattern/support"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	expected := []string{"Hoge(transformed)", "Fuga(transformed)"}

	cases := []struct {
		name    string
		cnvFunc ConvertFunc
	}{
		{name: "TemplateMethodPattern", cnvFunc: TemplateMethodPattern},
		{name: "StrategyPattern", cnvFunc: StrategyPattern},
	}
	for _, c := range cases {
		actual := make([]string, 0)

		c.cnvFunc(&actual)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Failed to convert by %v: expected=%v, actual=%v", c.name, expected, actual)
		}
	}
}

type ConvertFunc func(*[]string)

func TemplateMethodPattern(a *[]string) {
	for _, c := range NewConverterHolder(&support.Searcher{}).Converters {
		c.Do(c, a)
	}
}

func StrategyPattern(a *[]string) {
	ruleHolder := NewConvertRuleHolder()
	service := ConvertService{
		rules:    ruleHolder.rules,
		searcher: &support.Searcher{},
	}
	service.Do(a)
}
