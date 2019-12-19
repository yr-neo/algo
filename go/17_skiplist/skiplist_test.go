package _7_skiplist

import (
	"fmt"
	"testing"
)

func TestSkipList1(t *testing.T) {
	sl := NewSkipList()

	sl.Insert("leo", 95)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")

	sl.Insert("jack", 88)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")

	sl.Insert("lily", 100)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")

	t.Log(sl.Find("jack", 88))
	t.Log("-----------------------------")

	sl.Delete("leo", 95)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")
}

func TestSkipList2(t *testing.T) {
	sl := NewSkipList()
	fmt.Println(sl.PrintSummary())
	fmt.Println("===========================")

	sl.Insert(3, 3)
	sl.Insert(4, 4)
	sl.Insert(33, 33)
	sl.Insert(6, 6)
	sl.Insert(19, 19)
	sl.Insert(8, 8)
	sl.Insert(10, 10)
	sl.PrintDetail()
	fmt.Println(sl.PrintSummary())
	fmt.Println("===========================")

	sl.Delete(6, 6)
	sl.PrintDetail()
	fmt.Println(sl.PrintSummary())
	fmt.Println("===========================")

	fmt.Printf("find [6, 6]:%+v\n", sl.Find(6, 6))
}
