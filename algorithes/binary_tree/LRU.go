package binary_tree

// basic LRU
// type LRUCache struct {
//     cache []*entry
//     size int
// }

// type entry struct {
//     key int
//     value int
// }

// func Constructor(capacity int) LRUCache {
//     return LRUCache{
//         cache: make([]*entry, 0),
//         size: capacity,
//     }
// }

// func (this *LRUCache) Get(key int) int {
//     for i := len(this.cache) - 1; i >= 0; i-- {
//         if this.cache[i].key == key {
//             value := this.cache[i].value
//             this.MoveToHead(i)
//             return value
//         }
//     }

//     return -1
// }

// func (this *LRUCache) Put(key int, value int)  {
//    for i := len(this.cache) - 1; i >= 0; i-- {
//         if this.cache[i].key == key {
//             this.cache[i].value = value
//             this.MoveToHead(i)
//             return
//         }
//     }

//     if this.size <= len(this.cache) {
//         this.cache = this.cache[1:]
//     }
//     this.cache = append(this.cache, &entry {key, value})
// }

// func (this *LRUCache) MoveToHead(index int) {
//     for i := index + 1; i < len(this.cache); i++ {
//         this.cache[i], this.cache[i-1] = this.cache[i-1], this.cache[i]
//     }
// }

// high-level LRU
type LRUCache struct {
	size, capacity int
	cache          map[int]*DLinkedNode
	head, tail     *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	pre, next  *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}

	l.head.next = l.tail
	l.tail.pre = l.head

	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}

	node := this.cache[key]
	this.moveToHead(node)

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		if this.size >= this.capacity {
			removed := this.deleteFromTail()
			delete(this.cache, removed.key)
			this.size--
		}
		node := &DLinkedNode{key: key, value: value}
		this.cache[key] = node
		this.addToHead(node)
		this.size++
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.deleteNode(node)
	this.addToHead(node)
}

func (this *LRUCache) deleteFromTail() *DLinkedNode {
	node := this.tail.pre
	this.deleteNode(node)

	return node
}

func (this *LRUCache) deleteNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}
