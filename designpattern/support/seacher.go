package support

type Searcher struct {
}

func (s *Searcher) Do(name string) {
	println("search " + name + "...")
}
