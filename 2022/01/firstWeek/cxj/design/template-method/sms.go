package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Sms struct {
	notify Notify
}

func (s *Sms) genRandomCode(i int) string {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(1000000)))
	fmt.Println("code:", code)
	return code
}

func (s *Sms) sendVerifyCode(account, msg string, option ...interface{}) error {
	fmt.Println("Graceful sendVerifyCode")
	return nil
}

func (s *Sms) getCache(msg string) string {
	// get code redis db
	// return redis result
	fmt.Println("Graceful getCache")
	return "fedcba"
}

func (s *Sms) setCache(msg string) error {
	// set code to redis db
	fmt.Println("Graceful setCache")
	return nil
}

func (s *Sms) publicHooks() {
	fmt.Println("Graceful publicHooks")
	return
}
