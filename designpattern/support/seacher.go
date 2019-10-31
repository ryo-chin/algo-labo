package support

type Searcher struct {
}

func (s *Searcher) Do(name string) {
	println("searching by " + name + "...")
}
