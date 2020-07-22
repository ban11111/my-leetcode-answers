package skiplist

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()/3 + 516504610)
}

type SkipList struct {
	Head      Node
	MaxLayers int
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
	return opt[0] // 让他炸
	//return -1
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
				return nil // 最后一层也没有就找不到了
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
		if _, sub := next.(*SubItem); next == nil || sub || next.Value() > value { // 可以优化
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
	return target != nil
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
	current := sl.Head
	for ; current.Down() != nil; current = current.Down() {
	}
	i := 0
	for ; current.Next() != nil; i++ {
		current = current.Next()
	}
	return i
}
