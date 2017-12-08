package util


const NOT_FOUND int  = -1


type ArrayList struct {
	size int
	array []interface{}
}

func (list *ArrayList) Init() {
	list.size = 0
	list.array = make([]interface{}, 0, 0)
}

func (list *ArrayList) InitFromSlice(slice []interface{}) {
	list.array = slice
	list.size = len(slice)
}


func (list *ArrayList) Size() int {
	return list.size
}

func (list *ArrayList) IsEmpty() bool {
	return list.Size() == 0
}

func (list *ArrayList) Contains(object interface{}) bool {
	return list.IndexOf(object) != NOT_FOUND
}

func (list *ArrayList) Add(object interface{}) {
	list.array = append(list.array, object)
	list.size++
}

func (list *ArrayList) Remove(object interface{}) {
	index := list.IndexOf(object)
	if (index != NOT_FOUND) {
		list.RemoveByIndex(index)
	}
}

func (list *ArrayList) Clear() {
	list.Init()
}

func (list *ArrayList) Get(index int) interface{} {
	if list.isGoodIndex(index) {
		return list.array[index]
	} else {
		return nil
	}
}

func (list *ArrayList) Set(index int, element interface{}) {
	if list.isGoodIndex(index) {
		list.array[index] = element
	}
}

func (list *ArrayList) AddByIndex(index int, element interface{}) {
	list.array = append(list.array[:index], element, list.array[index+1])
	list.size++
}

func (list *ArrayList) RemoveByIndex(index int) {
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

func (list *ArrayList) IndexOf(object interface{}) int {
	for i, v := range list.array {
		if (v == object) {
			return i
		}
	}
	return NOT_FOUND
}

func (list *ArrayList) isGoodIndex(index int) bool {
	return (0 <= index) && (index <= list.size)
}
