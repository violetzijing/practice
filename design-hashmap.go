package main

type MyHashMap struct {
	Node *ListNode
	Next *ListNode
}

type ListNode struct {
	Key int
	Val int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		Node: &ListNode{},
		Next: nil,
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	if

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {

}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {

}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
