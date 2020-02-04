package Common

type Node struct {
	Key  int
	Val  int
	Last *Node
	Next *Node
}

type LRUCache struct {
	Cap  int
	Val  int
	Head *Node
	Hash map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Cap: capacity, Hash: make(map[int]*Node, capacity)}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.Hash[key]
	if !ok {
		return -1
	}

	if this.Head == val {
		return val.Val
	} else if this.Val == 2 {
		this.Head = val
		return val.Val
	} else if this.Head.Last == val {
		this.Head = val
		return val.Val
	}

	last := val.Last
	next := val.Next
	end := this.Head.Last
	last.Next = next
	next.Last = last
	val.Last = end
	val.Next = this.Head
	end.Next = val
	this.Head.Last = val
	this.Head = val
	return val.Val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Hash[key]; ok {
		this.Hash[key].Val = value
		this.Get(key)
		return
	}

	val := &Node{Key: key, Val: value}
	this.Val++
	if this.Head == nil {
		this.Head = val
		val.Next = val
		val.Last = val
		this.Hash[key] = val
		return
	}

	end := this.Head.Last
	if this.Val > this.Cap {
		this.Val--
		delete(this.Hash, end.Key)
		end = end.Last
	}
	end.Next = val
	val.Last = end
	val.Next = this.Head
	this.Head.Last = val
	this.Head = val
	this.Hash[key] = val
}
