package list

import (
	"cmp"

	"github.com/sword-demon/vtool/internal/list"
)

// LinkedList 双向链表
type LinkedList[T cmp.Ordered] = list.LinkedList[T]

// NewLinkedList 创建新的双向链表
func NewLinkedList[T cmp.Ordered]() *LinkedList[T] {
	return list.NewLinkedList[T]()
}
