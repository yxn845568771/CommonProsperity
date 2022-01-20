package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Email struct {
}

func NewEmail() INotify {
	return &Email{}
}

func (e *Email) genRandomCode(i int) string {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(1000000)))
	fmt.Println("code:", code)
	return code
}

func (e *Email) sendVerifyCode(account, msg string, option ...interface{}) error {
	fmt.Println("Graceful sendVerifyCode")
	return nil
}

func (e *Email) getCache(s string) string {
	// get code redis db
	// return redis result
	fmt.Println("Graceful getCache")
	return "abcdef"
}

func (e *Email) setCache(s string) error {
	// set code to redis db
	fmt.Println("Graceful setCache")
	return nil
}

func (e *Email) publicHooks() {
	fmt.Println("Graceful publicHooks")
	return
}
