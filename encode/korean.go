package encode

import "strings"

// Legacy multi-byte Korean encodings
// Charactor code: EUC-KR
func (e *Encode) EUC_KR() ([]byte, error) {
	b, err := encodeReader(strings.NewReader(e.Target), "euc-kr")
	return b, err
}
