package check

import "fmt"

func Err(op string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}
}
