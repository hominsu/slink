package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type CharSet string

const (
	NumberSet      CharSet = "0123456789"
	LowerCharSet   CharSet = "abcdefghijklmnopqrstuvwxyz"
	UpperCharSet   CharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SpecialCharSet CharSet = "~!@#$%^&*_+-=<>?"
	AllCharSet     CharSet = LowerCharSet + UpperCharSet + SpecialCharSet + NumberSet
)

var src = rand.NewSource(time.Now().UnixNano())

func RandString(n int, charSet ...CharSet) string {
	if len(charSet) == 0 {
		return randString(n, string(AllCharSet))
	}

	s := ""
	for _, set := range charSet {
		s += string(set)
	}
	return randString(n, s)
}

func randString(n int, charSet string) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(charSet) {
			sb.WriteByte(charSet[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
