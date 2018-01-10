package util

const PRESENT int = 7


type hashSet struct {
	hashMap map[interface{}]int
}

func newHashSet() *hashSet {
	hashSet := hashSet{}
	hashSet.Init()
	return &hashSet
}

func (set *hashSet) Init() {
	set.hashMap = make(map[interface{}]int)
}

func (set *hashSet) Size() int {
	return len(set.hashMap)
}

func (set *hashSet) IsEmpty() bool {
	return set.Size() == 0
}

func (set *hashSet) Contains(object interface{}) bool {
	return set.hashMap[object] == PRESENT
}

func (set *hashSet) Add(object interface{}) {
	set.hashMap[object] = PRESENT
}

func (set *hashSet) Remove(object interface{}) {
	delete(set.hashMap, object)
}

func (set *hashSet) Clear() {
	set.Init()
}
