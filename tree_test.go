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

func (t *TestItem) GetObject() interface{} {
	return t
}

func TestTree(t *testing.T) {
	// 6 -> 4 -> 2
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

func TestTreeEndlessLoop(t *testing.T) {
	// 2 -> 6 -> 4 -> 2
	// 5 -> 5
	list := []Row{
		&TestItem{ID: 1, Text: `one`, ParentID: 0},
		&TestItem{ID: 2, Text: `two`, ParentID: 4},
		&TestItem{ID: 3, Text: `three`, ParentID: 2},
		&TestItem{ID: 4, Text: `four`, ParentID: 6},
		&TestItem{ID: 5, Text: `five`, ParentID: 5},
		&TestItem{ID: 6, Text: `fix`, ParentID: 3},
		&TestItem{ID: 7, Text: `seven`, ParentID: 0},
	}
	data, _ := Build(list)
	Dump(data)
}

func TestTreeEndlessLoop2(t *testing.T) {
	// 2 -> 6 -> 2
	list := []Row{
		&TestItem{ID: 1, Text: `one`, ParentID: 0},
		&TestItem{ID: 2, Text: `two`, ParentID: 6},
		&TestItem{ID: 6, Text: `fix`, ParentID: 2},
	}
	data, _ := Build(list)
	Dump(data)
	//panic(``)
	//因为是通过自顶而下的方向构建树，死循环的结构没有顶部入口，故而会被忽略掉
	//                              (1)     (2)     (3)
	//                              / \      |       |
	//                           (10) (11)  (12)    (13)
	//                            /     \
	//                          (21)    (22)
	//                                    \
	//                                    (32)
}
