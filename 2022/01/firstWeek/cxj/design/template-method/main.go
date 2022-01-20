package main

func main() {
	sms := NewSms()
	sms.genRandomCode(3)
	email := NewEmail()
	email.genRandomCode(6)
}
