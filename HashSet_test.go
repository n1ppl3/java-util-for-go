package util

import "testing"

// size, isEmpty, contains, add, remove
func TestSet(t *testing.T) {
	set := HashSet{}
	set.Init()

	str := "myString"

	assertEquals(t, 0, set.Size())
	assertEquals(t, true, set.IsEmpty())
	assertEquals(t, false, set.Contains(str))

	set.Add(str)

	assertEquals(t, 1, set.Size())
	assertEquals(t, false, set.IsEmpty())
	assertEquals(t, true, set.Contains(str))

	set.Remove(str)
	assertEquals(t, false, set.Contains(str))
}

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Error("Expected ", expected, " got ", actual)
	}
}