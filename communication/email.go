package communication

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type SmtpInfo struct {
	Username   string
	Password   string
	SmtpServer string
	Port       int
}

func Email(info SmtpInfo, receivers []string, subject string, body string) {
	auth := smtp.PlainAuth("", info.Username, info.Password, info.SmtpServer)

	headers := fmt.Sprintf("Subject: %s\r\n", subject)
	headers += fmt.Sprintf("From: %s\r\n", info.Username)
	headers += fmt.Sprintf("To: %s\r\n", strings.Join(receivers, ","))
	headers += "\r\n"

	err := smtp.SendMail(info.SmtpServer+":25", auth, info.Username,
		receivers, []byte(headers+body))
	if err != nil {
		log.Fatal(err)
	}
}
