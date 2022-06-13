package main

import "fmt"

type MagicList struct {
	LastItem *Item
}

type Item struct {
	Value    int
	PrevItem *Item
}

func add(list *MagicList, value int) {
	item := Item{
		Value: value,
	}

	if list.LastItem == nil {
		list.LastItem = &item
	} else {
		item.PrevItem = list.LastItem
		list.LastItem = &item
	}
}

func toSlice(list *MagicList) []int {
	item := list.LastItem
	var result []int
	for ok := true; ok; ok = item != nil {
			result = append(result, item.Value)
			item = item.PrevItem
	}
	return result
}

func main() {
	magicList := &MagicList{}
	add(magicList, 25)
	add(magicList, 50)
	add(magicList, 75)

	fmt.Println(toSlice(magicList))
}