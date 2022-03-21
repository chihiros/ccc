package encode

import (
	"bytes"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

type Encode struct {
	Target string
}

// The code here cites this
// https://github.com/yuin/charsetutil/blob/8e6966c2a828ddc65c43daaf45cdc9610c903b1c/charsetutil.go#L145-L158
func encodeReader(s io.Reader, enc string) ([]byte, error) {
	e, _ := charset.Lookup(enc)
	if e == nil {
		return nil, fmt.Errorf("unsupported charset: %q", enc)
	}
	var buf bytes.Buffer
	writer := transform.NewWriter(&buf, e.NewEncoder())
	_, err := io.Copy(writer, s)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
