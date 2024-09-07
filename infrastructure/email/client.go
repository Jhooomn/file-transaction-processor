package email

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

type EmailService interface {
	Send(ctx context.Context, recipient, subject, body string) error
}

type emailService struct {
	from       string
	psw        string
	smtpServer string
}

func NewEmailService(from, psw, smtpServer string) EmailService {
	return &emailService{
		from:       from,
		psw:        psw,
		smtpServer: smtpServer,
	}
}

func (es *emailService) Send(ctx context.Context, recipient, subject, body string) error {
	host := strings.Split(es.smtpServer, ":")[0]
	auth := smtp.PlainAuth("", es.from, es.psw, host)

	client, err := smtp.Dial(es.smtpServer)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %v", err)
	}
	defer client.Close()

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	if err = client.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("failed to start TLS: %v", err)
	}

	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}

	if err = client.Mail(es.from); err != nil {
		return fmt.Errorf("failed to set sender email: %v", err)
	}
	if err = client.Rcpt(recipient); err != nil {
		return fmt.Errorf("failed to set recipient email: %v", err)
	}

	// Send the email message
	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to send email data: %v", err)
	}
	defer wc.Close()

	// Build the email message with headers
	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", es.from, recipient, subject, body)

	// Write the full message (headers + body)
	_, err = wc.Write([]byte(msg))
	if err != nil {
		return fmt.Errorf("failed to write email body: %v", err)
	}

	return nil
}
