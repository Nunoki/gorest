package array

// Filter returns only the items from the input array that satisfy the result of the compare
// function. This function behaves the same as Javascript's Array.filter() function.
//
// Example usage:
//
//	arr := []int{1, 2, 3, 4}
//	res := Filter(arr, func(item int) bool { return item > 2 })
//	// [3, 4]
func Filter[T any](input []T, compare func(T) bool) []T {
	filtered := make([]T, 0)
	for _, v := range input {
		if compare(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// Map applies the mapper transformation function to all items input the input array. This function
// behaves the same as Javascript's Array.map() function.
//
// Example usage:
//
//	arr := []int{1, 2, 3, 4}
//	res := Map(arr, func(item int) int { return item + 1 })
//	// [2, 3, 4, 5]
func Map[T any](input []T, mapper func(T) T) []T {
	for k, v := range input {
		input[k] = mapper(v)
	}
	return input
}

// Reduce applies the reducer function to all items from the input array, returning a single flat
// value as the output, which is of the same type as the items of the input. A reducer callback
// will receive the accumulator value, the value of the current item, and the index of the current
// item. This function behaves almost the same as Javascript's Array.reduce() function, except that
// the accumulator doesn't accept an initial value.
//
// Example usage:
//
//	arr := []int{1, 2, 3, 4}
//	res3 := array.Reduce(arr, func(acc int, item int, key int) int {
//	  if key > 2 {
//	    return acc
//	  } else {
//	    return acc + item
//	  }
//	})
//	// results in: 12
func Reduce[T any](input []T, reducer func(acc, item T, index int) T) (reduced T) {
	for k, v := range input {
		reduced = reducer(reduced, v, k)
	}
	return reduced
}

// Returns whether the array haystack contains the item needle.
//
// Example usage:
//
//	arr := []int{1, 2, 3, 4}
//	res := Contains(arr, 2)
//	// results in: true
func Contains[T comparable](haystack []T, needle T) bool {
	for _, i := range haystack {
		if i == needle {
			return true
		}
	}
	return false
}
