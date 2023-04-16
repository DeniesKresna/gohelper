package utstring

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

// n: total character
func RandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetEnv(key string, def ...string) string {
	resp := os.Getenv(key)
	if resp == "" {
		return Chain(def...)
	}
	return resp
}

func Chain(str ...string) string {
	for _, v := range str {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func ConvertToString(obj interface{}, def string) (res string) {
	switch v := obj.(type) {
	case int, int64, int32, float64, float32, bool:
		return fmt.Sprintf("%v", v)
	case string:
		return fmt.Sprintf("%v", v)
	default:
		return def
	}
}

func RemoveChars(str string, charsToRemove []string) string {
	for _, char := range charsToRemove {
		str = strings.ReplaceAll(str, char, "")
	}
	return str
}

func CamelToSnake(s string) string {
	symbols := []rune{'$', '%', '&', '@', '^', '[', ']', '{', '}', '(', ')'}

	// Remove any prefix that matches a symbol
	for len(s) > 0 && contains(symbols, rune(s[0])) {
		s = s[1:]
	}

	// Remove any suffix that matches a symbol
	for len(s) > 0 && contains(symbols, rune(s[len(s)-1])) {
		s = s[:len(s)-1]
	}

	var buf bytes.Buffer
	var prev rune
	for i, r := range s {
		if r == '.' {
			buf.WriteRune(r)
			prev = r
			continue
		}
		if unicode.IsSpace(r) {
			if prev != '_' && i > 0 {
				buf.WriteRune('_')
			}
			prev = '_'
			continue
		}
		if unicode.IsUpper(r) {
			if i > 0 && prev != '_' && !unicode.IsUpper(prev) {
				buf.WriteRune('_')
			}
			r = unicode.ToLower(r)
		}
		buf.WriteRune(r)
		prev = r
	}
	return strings.TrimSuffix(buf.String(), "_")
}

func contains(s []rune, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}
