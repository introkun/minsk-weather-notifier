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
    SmtpPort   int
    Receipents []string
}

func Email(info SmtpInfo, subject string, body string) {
    auth := smtp.PlainAuth("", info.Username, info.Password, info.SmtpServer)

    headers := fmt.Sprintf("Subject: %s\r\n", subject)
    headers += fmt.Sprintf("From: %s\r\n", info.Username)
    headers += fmt.Sprintf("To: %s\r\n", strings.Join(info.Receipents, ","))
    headers += "\r\n"

    err := smtp.SendMail(fmt.Sprintf("%s:%d", info.SmtpServer, info.SmtpPort), auth, info.Username,
        info.Receipents, []byte(headers+body))
    if err != nil {
        log.Fatal(err)
    }
}
