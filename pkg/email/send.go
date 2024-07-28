package email

import (
    "log"
    "net/smtp"
)

func SendEmail(to string, subject string, body string) {
    from := "your-email@gmail.com"
    password := "your-email-password"

    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject + "\n\n" +
        body

    err := smtp.SendMail("smtp.gmail.com:587",
        smtp.PlainAuth("", from, password, "smtp.gmail.com"),
        from, []string{to}, []byte(msg))

    if err != nil {
        log.Fatal(err)
    }
}
