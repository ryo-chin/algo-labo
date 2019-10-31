package designpattern

import (
	"github.com/ryo-chin/algo-labo/designpattern/support"
	"reflect"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	expected := []string{"Hoge(transformed)", "Fuga(transformed)"}
	actual := make([]string, 0)

	for _, c := range NewConverterHolder(&support.Searcher{}).Converters {
		c.Do(c, &actual)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Failed to convert: expected=%v, actual=%v", expected, actual)
	}
}
