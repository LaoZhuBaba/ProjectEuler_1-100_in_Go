package shared

// Return a function that will generate pentangular numbers with the input 1, -1, 2, -2, 3, -3, ...
// The returned values follow the series 1, 2, 5, 7, 12, 15, 22, ...
func PentGen() func() int {
	n := 1
	return func() int {
		// Using defer means that although n will evaluated before the return statement but the
		// value of n will not change until after the return value has been calculated.  Avoids
		// the need for a temporary variable to store the return value.
		defer func() {
			if n < 0 {
				n--
			}
			n *= -1
		}()
		return n * (3*n - 1) / 2
	}
}
