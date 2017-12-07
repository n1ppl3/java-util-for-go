package util

type Set interface {

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

	// boolean retainAll(Collection<?> c);

	// boolean removeAll(Collection<?> c);

	Clear()
}
