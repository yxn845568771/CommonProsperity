package main

func main() {
	sms := &Sms{}
	sms.genRandomCode(3)
	email := &Email{}
	email.genRandomCode(6)
}
