type MyCircularDeque struct {
     len int 	//当前长度
     cap int 	//总容量
     data []int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        len: 0,
        cap: k,
        data: make([]int, 0, 10),
    }
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    this.data = append([]int{value}, this.data...)
    this.len ++
    return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
     if this.IsFull() {
        return false
    }
    this.data = append(this.data, value)
    this.len ++
    return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    this.data = append(this.data[:0], this.data[1:]...)
    this.len --
    return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    this.data = this.data[0:this.len-1]
    this.len --
    return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.len <= 0 {
        return -1
    }
    return this.data[0]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.len <= 0 {
        return -1
    }
    return this.data[this.len-1]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    if this.len == 0 {
        return true
    }
    return false
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    if this.len >= this.cap {
        return true
    }
    return false
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */