package token

import (
	"io"
	"os"
	"strings"
	"unicode"
)

func Tokenizer(f *os.File) []string {
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	text := string(data)
	text = strings.ToLower(text)
	token :=strings.FieldsFunc(text,func(r rune)bool{
		return !unicode.IsLetter(r)&&!unicode.IsNumber(r)
	})
	return token
}
