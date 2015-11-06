package main

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
    b := make([]byte, n)
    maxRange := len(letterBytes)
    for i := range b {
    	rand.Seed(time.Now().UnixNano())
        b[i] = letterBytes[rand.Intn(maxRange)]
    }
    return string(b)
}

func getRandomString() string {
	rand.Seed(time.Now().Unix())

	maxLength := 4
	strLen := rand.Intn(maxLength - 1) + 1
	return randStringBytes(strLen)
}
