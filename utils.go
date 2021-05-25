package tree

import (
	"encoding/json"
	"fmt"
)

type Row interface {
	GetID() int64
	GetParentID() int64
}

func Build(list []Row) (*Tree, error) {
	data := New()
	parents := map[int64][]Row{}
	for _, row := range list {
		if row.GetParentID() == 0 {
			item := NewItem(row.GetID(), row)
			data.Add(item)
			continue
		}
		if _, ok := parents[row.GetParentID()]; !ok {
			parents[row.GetParentID()] = []Row{}
		}
		parents[row.GetParentID()] = append(parents[row.GetParentID()], row)
	}
	err := genChildrenItem(data, parents)
	return data, err
}

func genChildrenItem(data *Tree, parents map[int64][]Row) error {
	return data.Range(func(item *Item) error {
		children, ok := parents[item.GetID()]
		if !ok {
			return nil
		}
		items := make([]*Item, len(children))
		for index, child := range children {
			itemNew := NewItem(child.GetID(), child)
			items[index] = itemNew
		}
		item.Add(items...)
		return genChildrenItem(item.Children, parents)
	})
}

// Dump 输出对象和数组的结构信息
func Dump(m interface{}, args ...bool) (r string) {
	v, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	r = string(v)
	l := len(args)
	if l < 1 || args[0] {
		fmt.Println(r)
	}
	return
}
