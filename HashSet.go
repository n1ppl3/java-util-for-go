package util

const PRESENT int = 7


type HashSet struct {
	hashMap map[interface{}]int
}

func (set *HashSet) Init() {
	set.hashMap = make(map[interface{}]int)
}

func (set *HashSet) Size() int {
	return len(set.hashMap)
}

func (set *HashSet) IsEmpty() bool {
	return set.Size() == 0
}

func (set *HashSet) Contains(object interface{}) bool {
	return set.hashMap[object] == PRESENT
}

func (set *HashSet) Add(object interface{}) {
	set.hashMap[object] = PRESENT
}

func (set *HashSet) Remove(object interface{}) {
	delete(set.hashMap, object)
}

func (set *HashSet) Clear() {
	set.Init()
}
