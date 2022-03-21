package ccc

import "github.com/chihiros/change-character-code/encode"

func Encode(str string) *encode.Encode {
	return &encode.Encode{
		Target: str,
	}
}
