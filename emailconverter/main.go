package main

import (
	"fmt"
)

func main() {
	recipientName := "John Doe"
	recipientEmail := "john@example.com"
	subtle := "Welcome to Our Service"
	message := "Hello John Doe"
	messageBody := "Thank you for joining our service! We are excited to have you on board."

	result := emailGenerater(recipientName,
		recipientEmail,
		subtle,
		message, messageBody)
	fmt.Println(result)
}

func emailGenerater(recipientName, recipientEmail, subject, message, messageBody string) (email string) {
	email = fmt.Sprintf(
		`From: no-reply@example.com
To: %s <%s>
Subject: %s,


%s


%s

Best regards,
Your Company Team`,
		recipientName,
		recipientEmail,
		subject, message,
		messageBody)

	return
}
