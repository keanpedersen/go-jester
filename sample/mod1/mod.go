package mod1

import "fmt"

func FizzBuzz(max int) (ret []string) {

	for i := 1; i <= max; i++ {
		line := ""
		if i%3 == 0 {
			line += "fizz"
		}
		if i%5 == 0 {
			line += "buzz"
		}

		if i%3 != 0 && i%5 != 0 {
			line += fmt.Sprintf("%d", i)
		}
		ret = append(ret, line)
	}
	return ret
}
