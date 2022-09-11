package util

import "math/rand"

func RandomName(n int) (data string) {
	var str = "asdfghjklqwertyuiopzxcvbnmZXCVBNMASDFGHJKLQWERTYUIOP"
	result := make([]byte, n)
	for i := range result {
		result[i] = str[rand.Intn(len(str))]
	}
	data = string(result)
	return
}
