package handler
import (
    "fmt"
    "math/rand"
    "time"
    "gopkg.in/gomail.v2"

)


func generateOTP() string {
    rand.Seed(time.Now().UnixNano())
    otp := ""
    for i := 0; i < 4; i++ {
        otp += fmt.Sprintf("%d", rand.Intn(10))
    }
    return otp
}
func sendOTPEmail(to, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@example.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s", otp))

	d := gomail.NewDialer("smtp.example.com", 587, "your-email@example.com", "your-email-password")

	// Send the email
	return d.DialAndSend(m)
}
