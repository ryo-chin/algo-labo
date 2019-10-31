package designpattern

import "github.com/ryo-chin/algo-labo/designpattern/support"

//
// Base
//
type ConvertRule interface {
	Search(s *support.Searcher) string
	Transform(string) string
}

//
// Concrete
//
type HogeConvertRule struct {
}

func (r *HogeConvertRule) Search(s *support.Searcher) string {
	name := "Hoge"
	s.Do(name)
	return name
}

func (r *HogeConvertRule) Transform(name string) string {
	return name + "(transformed)"
}

type FugaConvertRule struct {
}

func (r *FugaConvertRule) Search(s *support.Searcher) string {
	name := "Fuga"
	s.Do(name)
	return name
}

func (r *FugaConvertRule) Transform(name string) string {
	return name + "(transformed)"
}

type ConvertService struct {
	rules    []ConvertRule
	searcher *support.Searcher
}

func (s *ConvertService) Do(inputs *[]string) {
	for _, r := range s.rules {
		name := r.Search(s.searcher)
		*inputs = append(*inputs, r.Transform(name))
	}
}

//
// Holder
//
type ConvertRuleHolder struct {
	rules []ConvertRule
}

func NewConvertRuleHolder() ConvertRuleHolder {
	return ConvertRuleHolder{
		[]ConvertRule{
			&HogeConvertRule{},
			&FugaConvertRule{},
		},
	}

}
