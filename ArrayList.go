package util

const NOT_FOUND int  = -1


type arrayList struct {
	size int
	array []interface{}
}

func NewArrayList() *arrayList {
	list := arrayList{}
	list.Init()
	return &list
}

func (list *arrayList) Init() {
	list.size = 0
	list.array = make([]interface{}, 0, 0)
}

func InitFromSlice(slice []interface{}) *arrayList {
	list := arrayList{}
	list.size = len(slice)
	list.array = slice
	return &list
}


func (list *arrayList) Size() int {
	return list.size
}

func (list *arrayList) IsEmpty() bool {
	return list.Size() == 0
}

func (list *arrayList) Contains(object interface{}) bool {
	return list.IndexOf(object) != NOT_FOUND
}

func (list *arrayList) Add(object interface{}) {
	list.array = append(list.array, object)
	list.size++
}

func (list *arrayList) Remove(object interface{}) {
	index := list.IndexOf(object)
	if (index != NOT_FOUND) {
		list.RemoveByIndex(index)
	}
}

func (list *arrayList) Clear() {
	list.Init()
}

func (list *arrayList) Get(index int) interface{} {
	if list.isGoodIndex(index) {
		return list.array[index]
	} else {
		return nil
	}
}

func (list *arrayList) Set(index int, element interface{}) {
	if list.isGoodIndex(index) {
		list.array[index] = element
	}
}

func (list *arrayList) AddByIndex(index int, element interface{}) {
	list.array = append(list.array[:index], element, list.array[index+1])
	list.size++
}

func (list *arrayList) RemoveByIndex(index int) {
	if list.isGoodIndex(index) {
		isFirst := (index == 0)
		isLast := (index == list.size-1)
		if isFirst {
			list.array = append(list.array[index+1:])
		} else if isLast {
			list.array = append(list.array[:index])
		} else {
			list.array = append(list.array[:index], list.array[index+1])
		}
		list.size--
	}
}

func (list *arrayList) IndexOf(object interface{}) int {
	for i, v := range list.array {
		if (v == object) {
			return i
		}
	}
	return NOT_FOUND
}

func (list *arrayList) isGoodIndex(index int) bool {
	return (0 <= index) && (index <= list.size)
}
