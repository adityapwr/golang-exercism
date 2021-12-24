package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	for i, v := range slice {
		if i == index {
			return v, true
		}
	}
	return 0, false
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	_, ok := GetItem(slice, index)
	if ok {
		slice[index] = value
	} else {
		slice = append(slice, value)
	}
	return slice

}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	var output []int
	for i := 0; i < length; i++ {
		output = append(output, value)
	}
	return output
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	_, ok := GetItem(slice, index)
	if !ok {
		return slice
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}
