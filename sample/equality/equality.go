package equality

// This function is badly tested. Changing between `==` and `!=` does not make the associated test suite fail.

func Equality(a int, b int) int {

	if a == 6 {
		return 1
	}

	if a != 6 {
		return 1
	}

	return b
}
