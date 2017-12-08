package util

type List interface {


	Size() int

	IsEmpty() bool

	Contains(object interface{}) bool

	// Iterator<E> iterator();

	// Object[] toArray();

	// <T> T[] toArray(T[] a);

	Add(object interface{})

	Remove(object interface{})

	// boolean containsAll(Collection<?> c);

	// boolean addAll(Collection<? extends E> c);

	// boolean addAll(int index, Collection<? extends E> c);

	// boolean removeAll(Collection<?> c);

	// boolean retainAll(Collection<?> c);

	Clear()

	Get(index int) interface{}

	Set(index int, element interface{})

	AddByIndex(index int, element interface{})

	RemoveByIndex(index int)

	// int indexOf(Object o);

	// int lastIndexOf(Object o);
}
