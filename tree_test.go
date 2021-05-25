package tree

import (
	"testing"
)

type TestItem struct {
	ID       int64  `json:"id"`
	Text     string `json:"text"`
	Icon     string `json:"icon,omitempty"`
	ParentID int64  `json:"parentID"`
}

func (t *TestItem) GetID() int64 {
	return t.ID
}

func (t *TestItem) GetParentID() int64 {
	return t.ParentID
}

func TestTree(t *testing.T) {
	list := []Row{
		&TestItem{ID: 1, Text: `one`, ParentID: 0},
		&TestItem{ID: 2, Text: `two`, ParentID: 4},
		&TestItem{ID: 3, Text: `three`, ParentID: 0},
		&TestItem{ID: 4, Text: `four`, ParentID: 6},
		&TestItem{ID: 5, Text: `five`, ParentID: 0},
		&TestItem{ID: 6, Text: `fix`, ParentID: 0},
		&TestItem{ID: 7, Text: `seven`, ParentID: 0},
	}
	data, _ := Build(list)
	Dump(data)
	if (*data)[3].Row.(*TestItem).ID != 6 {
		t.Fatalf(`(*data)[3].Row.(*TestItem).ID == %v, expected 6`, (*data)[3].Row.(*TestItem).ID)
	}
	if (*(*data)[3].Children)[0].Row.(*TestItem).ID != 4 {
		t.Fatalf(`(*(*data)[3].Children)[0].Row.(*TestItem).ID == %v, expected 4`, (*(*data)[3].Children)[0].Row.(*TestItem).ID)
	}
	if (*(*(*data)[3].Children)[0].Children)[0].Row.(*TestItem).ID != 2 {
		t.Fatalf(`(*(*(*data)[3].Children)[0].Children)[0].Row.(*TestItem).ID == %v, expected 2`, (*(*(*data)[3].Children)[0].Children)[0].Row.(*TestItem).ID)
	}
}
