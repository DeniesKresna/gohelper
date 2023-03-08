package utstring

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
