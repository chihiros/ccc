package ccc

import "ccc/encode"

func Encode(str string) *encode.Encode {
	return &encode.Encode{
		Target: str,
	}
}
