package charsetsetter

import (
	"io"

	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
)

func Charset(currcharset string, input io.Reader) (io.Reader, error) {
	return charset.NewReader(currcharset, input)
}
