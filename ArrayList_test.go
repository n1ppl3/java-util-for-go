package util

import (
	"testing"
	"fmt"
)

func TestList(t *testing.T) {
	var arrayList ArrayList
	arrayList = ArrayList{}

	arrayList.Init()
	fmt.Println("arrayList:", arrayList)

	str := "myString"

	assertEquals(t, 0, arrayList.Size())
	assertEquals(t, true, arrayList.IsEmpty())
	assertEquals(t, false, arrayList.Contains(str))

	arrayList.Add(str)
	fmt.Println("arrayList:", arrayList)

	assertEquals(t, 1, arrayList.Size())
	assertEquals(t, false, arrayList.IsEmpty())
	assertEquals(t, true, arrayList.Contains(str))

	arrayList.Remove(str)
	assertEquals(t, false, arrayList.Contains(str))
}


func TestRemoveByIndexOneSizeArray(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()
	fmt.Println("arrayList:", arrayList)

	arrayList.Add(18950)
	fmt.Println("arrayList:", arrayList)

	arrayList.RemoveByIndex(0)
	fmt.Println("arrayList:", arrayList)
}

func TestRemoveByIndexFirstElement(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()
	fmt.Println("arrayList:", arrayList)

	arrayList.Add(1)
	arrayList.Add(2)
	arrayList.Add(3)
	arrayList.Add(4)
	fmt.Println("arrayList:", arrayList)

	arrayList.RemoveByIndex(0)
	fmt.Println("arrayList:", arrayList)
}

func TestRemoveByIndexLastElement(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()
	fmt.Println("arrayList:", arrayList)

	arrayList.Add(1)
	arrayList.Add(2)
	arrayList.Add(3)
	arrayList.Add(4)
	fmt.Println("arrayList:", arrayList)

	arrayList.RemoveByIndex(3)
	fmt.Println("arrayList:", arrayList)
}
