package main

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// false
	return true
}

func (this *Iterator) next() int {
	// false
	return 0
}

type PeekingIterator struct {
	iter     *Iterator
	_next    int
	_hasNext bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:     iter,
		_next:    iter.next(),
		_hasNext: iter.hasNext(),
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this._hasNext
}

func (this *PeekingIterator) next() int {
	ret := this._next
	this._hasNext = this.iter.hasNext()
	if this._hasNext {
		this._next = this.iter.next()
	}

	return ret
}

func (this *PeekingIterator) peek() int {
	return this._next
}
