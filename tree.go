package tree

type Tree []*Item

func New() *Tree {
	return &Tree{}
}

func (t *Tree) Add(item ...*Item) *Tree {
	*t = append(*t, item...)
	return t
}

func (t *Tree) Size() int {
	return len(*t)
}

func (t *Tree) Range(fn func(item *Item) error) error {
	var err error
	for _, item := range *t {
		err = fn(item)
		if err != nil {
			break
		}
	}
	return err
}

func (t *Tree) Search(searcher func(item *Item) bool) *Item {
	for _, item := range *t {
		if found := item.Search(searcher); found != nil {
			return found
		}
	}
	return nil
}

func (t *Tree) Reset() *Tree {
	*t = (*t)[0:0]
	return t
}
