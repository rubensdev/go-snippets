package datastructures

type MyCollection struct {
	Data []int
}

func (c *MyCollection) Iterate() *MyCollectionIterator {
	return &MyCollectionIterator{c, 0}
}

type MyCollectionIterator struct {
	collection *MyCollection
	index      int
}

func (it *MyCollectionIterator) Next() bool {
	if it.index >= len(it.collection.Data) {
		return false
	}
	it.index++
	return true
}

func (it *MyCollectionIterator) Value() int {
	return it.collection.Data[it.index-1]
}
