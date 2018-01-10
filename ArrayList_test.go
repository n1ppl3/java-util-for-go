package util

import (
	"testing"
	"fmt"
)

func TestList(t *testing.T) {
	var list List
	list = NewArrayList()

	fmt.Println("list:", list)

	str := "myString"

	assertEquals(t, 0, list.Size())
	assertEquals(t, true, list.IsEmpty())
	assertEquals(t, false, list.Contains(str))

	list.Add(str)
	fmt.Println("list:", list)

	assertEquals(t, 1, list.Size())
	assertEquals(t, false, list.IsEmpty())
	assertEquals(t, true, list.Contains(str))

	list.Remove(str)
	assertEquals(t, false, list.Contains(str))
}


func TestRemoveByIndexOneSizeArray(t *testing.T) {
	arrayList := NewArrayList()
	fmt.Println("arrayList:", arrayList)

	arrayList.Add(18950)
	fmt.Println("arrayList:", arrayList)

	arrayList.RemoveByIndex(0)
	fmt.Println("arrayList:", arrayList)
}

func TestRemoveByIndexFirstElement(t *testing.T) {
	arrayList := InitFromSlice(IntSliceToInterfaceSlice([]int{1,2,3,4}))
	fmt.Println("arrayList:", arrayList)

	arrayList.RemoveByIndex(0)
	fmt.Println("arrayList:", arrayList)
}

func TestRemoveByIndexLastElement(t *testing.T) {
	arrayList := InitFromSlice(IntSliceToInterfaceSlice([]int{1,2,3,4}))
	fmt.Println("arrayList:", arrayList)

	arrayList.RemoveByIndex(3)
	fmt.Println("arrayList:", arrayList)
}
