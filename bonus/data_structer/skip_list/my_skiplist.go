package skiplist

import (
	"math/rand"
	"time"
)

// this is a one-way linklist version, might do a two-way linklist version later which conducts higher performance for deleting

func init() {
	rand.Seed(time.Now().UnixNano()/3 + 516504610)
}

type SkipList struct {
	Head      Node
	MaxLayers int
	length    int // number of all items
}

func NewSkipList(maxLayers int) *SkipList {
	head := &Item{}
	down := new(SubItem)
	head.down = down
	for i := 0; i < maxLayers-2; i++ {
		down.down = new(SubItem)
		down = down.down.(*SubItem)
	}
	return &SkipList{
		Head:      head,
		MaxLayers: maxLayers,
	}
}

// There are 2 kinds of nodes, item and sub-item, item contains value while sub-item doesn't
// item/sub-item stores pointer of lower sub-item and the pointer of the next node, simple illustration as bellow:
//
//         item --------------------------------> item ----> nil
//         sub ------> item --------------------> sub -----> nil
//         sub ------> sub -------> item -------> sub -----> nil
//         sub ------> sub -------> sub --------> sub -----> nil

type Node interface {
	Value(opt ...int) int
	Next() Node
	Down() Node
	//
	//SetValue(v int)
	SetNext(Node)
	SetDown(Node)
}

type Item struct {
	value int
	next  Node
	down  Node
}

func (i *Item) Value(opt ...int) int {
	return i.value
}

func (i *Item) Next() Node {
	return i.next
}

func (i *Item) Down() Node {
	return i.down
}

func (i *Item) SetValue(v int) {
	i.value = v
}

func (i *Item) SetNext(n Node) {
	i.next = n
}

func (i *Item) SetDown(n Node) {
	i.down = n
}

type SubItem struct {
	next Node
	down Node
}

func (s *SubItem) Value(opt ...int) int {
	return opt[0] // will panic; panic("wont come to this point, or your program is logically wrong")
}

func (s *SubItem) Next() Node {
	return s.next
}

func (s *SubItem) Down() Node {
	return s.down
}

func (s *SubItem) SetValue(v int) {
}

func (s *SubItem) SetNext(n Node) {
	s.next = n
}

func (s *SubItem) SetDown(n Node) {
	s.down = n
}

//func (sl *SkipList) reachingDown(value int) Node {
//
//}

func (sl *SkipList) find(value int) Node {
	for current := sl.Head; ; {
		next := current.Next()
		if _, sub := next.(*SubItem); next == nil || sub || next.Value() > value {
			if current.Down() == nil {
				return nil // when it goes beyond the lowest level, you won't be able to find what you need
			}
			current = current.Down()
			continue
		}
		if next.Value() == value {
			return next
		}
		if next.Value() < value {
			current = next
			continue
		}
	}
}

func (sl *SkipList) insert(value int) bool {
	layers := sl.random()
	updateNodes := make([]Node, layers)
	current := sl.Head
	for depth, index := 0, 0; current != nil; {
		next := current.Next()
		if _, sub := next.(*SubItem); next == nil || sub || next.Value() > value { // optimizable
			if depth >= sl.MaxLayers-layers {
				updateNodes[index] = current
				index++
			}
			current = current.Down()
			depth++
			continue
		}
		if next.Value() == value {
			return false
		}
		if next.Value() < value {
			current = next
			continue
		}
	}

	for i, n := range updateNodes {
		if i == 0 {
			n.SetNext(&Item{value: value, next: n.Next()})
		} else {
			downer := &SubItem{next: n.Next()}
			updateNodes[i-1].Next().SetDown(downer)
			n.SetNext(downer)
		}
	}
	sl.length++
	return true
}

func (sl *SkipList) delete(value int) bool {
	var target Node
	current := sl.Head
	for ; current != nil; {
		next := current.Next()
		if _, sub := next.(*SubItem); next == nil || sub || next.Value() > value {
			current = current.Down()
			continue
		}
		if next.Value() == value {
			target = next
			current.SetNext(next.Next())
			current = current.Down()
			break
		}
		if next.Value() < value {
			current = next
			continue
		}
	}
	for ; current != nil; {
		next := current.Next()
		if next == nil || next != target.Down() {
			current = next
			continue
		}
		if next == target.Down() {
			target = next
			current.SetNext(next.Next())
			current = current.Down()
			continue
		}
	}
	if target != nil {
		sl.length--
		return true
	}
	return false
}

func (sl *SkipList) random() int {
	layer := 1
	for i := 1; i < sl.MaxLayers; i++ {
		if rand.Int63()%2 == 1 {
			layer++
		} else {
			return layer
		}
	}
	return layer
}

func (sl *SkipList) len() int {
	return sl.length
}

func (sl *SkipList) deepLen() int {
	current := sl.Head
	for ; current.Down() != nil; current = current.Down() {
	}
	i := 0
	for ; current.Next() != nil; current = current.Next() {
		i++
	}
	return i
}
