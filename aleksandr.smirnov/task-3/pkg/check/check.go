package check

import (
	"fmt"
	"io"
)

func Err(op string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}
}

func Close(name string, c io.Closer) {
	Err("close "+name, c.Close())
}
