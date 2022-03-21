package encode

import "strings"

func (e *Encode) ShiftJIS() ([]byte, error) {
	b, err := encodeReader(strings.NewReader(e.Target), "sjis")
	return b, err
}

func (e *Encode) Windows31J() ([]byte, error) {
	b, err := encodeReader(strings.NewReader(e.Target), "Windows-31J")
	return b, err
}
