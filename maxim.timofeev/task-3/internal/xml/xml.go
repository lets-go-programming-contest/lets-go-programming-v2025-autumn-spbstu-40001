package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func NewDecoder(r io.Reader) *xml.Decoder {
	decoder := xml.NewDecoder(r)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		cs := strings.ToLower(strings.TrimSpace(charset))
		switch cs {
		case "utf-8", "utf8":
			return input, nil
		case "windows-1251", "cp1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		case "koi8-r":
			return charmap.KOI8R.NewDecoder().Reader(input), nil
		case "iso-8859-5":
			return charmap.ISO8859_5.NewDecoder().Reader(input), nil
		default:
			return input, nil
		}
	}

	return decoder
}

func ReadXML(path string, out any) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read XML file: %w", err)
	}

	decoder := NewDecoder(bytes.NewReader(file))

	if err := decoder.Decode(out); err != nil {
		return fmt.Errorf("decode XML: %w", err)
	}

	return nil
}
