package ccc

import "github.com/chihiros/ccc/encode"

func Encode(str string) *encode.Encode {
	return &encode.Encode{
		Target: str,
	}
}
