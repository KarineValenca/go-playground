package main

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"io"
)

func main() {
	c := getCode("test@example.com")
	fmt.Println(c)
	c = getCode("text@exampl.com")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("secretkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}