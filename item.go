package tree

func NewItem(id int64, row Row, items ...*Item) *Item {
	return &Item{
		Row:      row,
		Children: New().Add(items...),
	}
}

type Item struct {
	Row      `json:"row,omitempty"`
	Children *Tree `json:"children"`
}

func (a *Item) Add(item ...*Item) *Item {
	a.Children.Add(item...)
	return a
}

func (a *Item) Search(searcher func(item *Item) bool) *Item {
	if searcher(a) {
		return a
	}
	return a.Children.Search(searcher)
}
