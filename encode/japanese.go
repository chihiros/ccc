package encode

import "strings"

func (e *Encode) ShiftJIS() ([]byte, error) {
	b, err := encodeReader(strings.NewReader(e.Target), "sjis")
	return b, err
}

func (e *Encode) EUC_JP() ([]byte, error) {
	b, err := encodeReader(strings.NewReader(e.Target), "euc-jp")
	return b, err
}

func (e *Encode) ISO_2022_JP() ([]byte, error) {
	b, err := encodeReader(strings.NewReader(e.Target), "iso-2022-jp")
	return b, err
}
