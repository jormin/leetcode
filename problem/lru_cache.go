package problem

import (
	"fmt"
	"strings"
)

// 运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。
// 实现 LRUCache 类：
//
// LRUCache(int capacity) 以正整数作为容量capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value)如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
// 进阶：你是否可以在O(1) 时间复杂度内完成这两种操作？
//
// 示例：
// 输入
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// 输出
// [null, null, null, 1, null, -1, null, -1, 3, 4]
//
// 解释
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // 缓存是 {1=1}
// lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
// lRUCache.get(1);    // 返回 1
// lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
// lRUCache.get(2);    // 返回 -1 (未找到)
// lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
// lRUCache.get(1);    // 返回 -1 (未找到)
// lRUCache.get(3);    // 返回 3
// lRUCache.get(4);    // 返回 4
//
// 提示：
//
// 1 <= capacity <= 3000
// 0 <= key <= 10000
// 0 <= value <= 105
// 最多调用 2 * 105 次 get 和 put
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/lru-cache
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// Constructor 构造
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		head:     &DoubleLinkNode{},
		tail:     &DoubleLinkNode{},
		cache:    map[int]*DoubleLinkNode{},
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

// LRUCache LRU缓存
type LRUCache struct {
	size       int
	capacity   int
	head, tail *DoubleLinkNode // 虚拟头尾节点
	cache      map[int]*DoubleLinkNode
}

// Get 获取
func (l *LRUCache) Get(key int) int {
	node, ok := l.cache[key]
	if !ok {
		return -1
	}
	// 查询到之后需要将该节点移动到头节点
	l.moveToHead(node)
	return node.val
}

// Put 存储
func (l *LRUCache) Put(key int, value int) {
	// 节点存在
	if l.Get(key) != -1 {
		node := l.cache[key]
		node.val = value
		l.moveToHead(node)
		return
	}
	// 节点不存在
	node := &DoubleLinkNode{
		key: key,
		val: value,
	}
	l.addToHead(node)
	l.cache[key] = node
	l.size++
	// 如果size大于capacity
	if l.size > l.capacity {
		node := l.removeTail()
		delete(l.cache, node.key)
		l.size--
	}
}

// removeNode 移除节点
func (l *LRUCache) removeNode(node *DoubleLinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

// addToHead 添加到头节点
func (l *LRUCache) addToHead(node *DoubleLinkNode) {
	node.next = l.head.next
	node.pre = l.head
	l.head.next.pre = node
	l.head.next = node
}

// moveToHead 移动节点到头节点
func (l *LRUCache) moveToHead(node *DoubleLinkNode) {
	l.removeNode(node)
	l.addToHead(node)
}

// removeTail 移除尾节点
func (l *LRUCache) removeTail() *DoubleLinkNode {
	node := l.tail.pre
	l.removeNode(node)
	return node
}

// String 字符串
func (l *LRUCache) String() string {
	var arr []string
	node := l.head
	for node != nil {
		arr = append(arr, fmt.Sprintf("%d", node.val))
		node = node.next
	}
	return strings.Join(arr, ",")
}

// DoubleLinkNode 双向链表
type DoubleLinkNode struct {
	key, val  int
	pre, next *DoubleLinkNode
}
