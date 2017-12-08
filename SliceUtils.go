package util

/*
https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces-in-go
*/
func IntSliceToInterfaceSlice(data []int) []interface{} {
	interfaceSlice := make([]interface{}, len(data))
	for i, d := range data {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
