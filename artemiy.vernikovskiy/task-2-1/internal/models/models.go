package models

import (
	"io"
)

type Wrapper struct {
	Cin  io.Reader
	Cout io.Writer
}
