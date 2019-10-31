package designpattern

import "github.com/ryo-chin/algo-labo/designpattern/support"

//
// Base
//
type Converter interface {
	Do(c Converter, inputs *[]string)
	search() string
	transform(string) string
}

type abstractConverter struct {
	searcher *support.Searcher
}

func (_c *abstractConverter) Do(c Converter, inputs *[]string) {
	res := c.search()
	*inputs = append(*inputs, c.transform(res))
}

//
// Concrete
//
type HogeConverter struct {
	*abstractConverter
}

func (c *HogeConverter) search() string {
	c.searcher.Do("Hoge")
	return "Hoge"
}

func (c *HogeConverter) transform(name string) string {
	println("transform to HogeConverter")
	return name + "(transformed)"
}

type FugaConverter struct {
	*abstractConverter
}

func (c *FugaConverter) search() string {
	c.searcher.Do("Fuga")
	return "Fuga"
}
func (c *FugaConverter) transform(name string) string {
	println("transform to FugaConverter")
	return name + "(transformed)"
}

//
// Holder
//
func NewConverterHolder(searcher *support.Searcher) ConverterHolder {
	ac := &abstractConverter{searcher: searcher}
	return ConverterHolder{[]Converter{
		&HogeConverter{ac},
		&FugaConverter{ac},
	}}
}

type ConverterHolder struct {
	Converters []Converter
}
