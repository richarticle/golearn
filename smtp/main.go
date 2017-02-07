package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

func main() {

	s := NewSmtp("smtp.server.com", 25, "user@server.com", "PASSWORD")
	err := s.SendMail("user1@gmail.com,user2@gmail.com", "Subject", "Body")
	if err != nil {
		fmt.Println(err)
	}
}

// Smtp configuration
type Smtp struct {
	Server   string
	Port     uint16
	From     string
	Password string
	Auth     smtp.Auth
}

// NewSmtp creates a SMTP instance
func NewSmtp(server string, port uint16, from, password string) *Smtp {
	s := &Smtp{
		Server:   server,
		Port:     port,
		From:     from,
		Password: password}

	s.Auth = smtp.PlainAuth("", from, password, server)

	return s
}

// SendMail sends out a mail
func (s *Smtp) SendMail(to, subject, body string) error {

	tos := strings.Split(to, ",")

	header := make(map[string]string)
	header["From"] = s.From
	header["To"] = to
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	msg := &bytes.Buffer{}
	for k, v := range header {
		fmt.Fprintf(msg, "%s: %s\r\n", k, v)
	}
	fmt.Fprintf(msg, "\r\n%s", base64.StdEncoding.EncodeToString([]byte(body)))

	return sendMail(fmt.Sprintf("%s:%d", s.Server, s.Port), s.Auth, s.From, tos, msg.Bytes())
}

// sendMail implement SMTP send function
func sendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if ok, _ := c.Extension("STARTTLS"); ok {
		host, _, _ := net.SplitHostPort(addr)
		config := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         host,
		}
		if err = c.StartTLS(config); err != nil {
			return err
		}
	}

	if ok, _ := c.Extension("AUTH"); ok {
		if err = c.Auth(a); err != nil {
			return err
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
