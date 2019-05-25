package providers

import (
	"crypto/tls"
	"fmt"
	"github.com/hirsch88/go-trophy-server/config"
	"go.uber.org/zap"
	"net"
	"net/mail"
	"net/smtp"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

type smtpMailProvider struct {
	log    *zap.SugaredLogger
	config *config.MailConfig
}

func (p *smtpMailProvider) Send(to string, subject string, body string) error {
	from := mail.Address{p.config.Name, p.config.From}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += MIME + "\r\n" + body

	// Connect to the SMTP Server
	servername := p.config.Host + ":" + p.config.Port
	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", p.config.From, p.config.Password, host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		p.log.Error(err)
		return err
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		p.log.Error(err)
		return err
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		p.log.Error(err)
		return err
	}

	// step 2: add all from and to
	if err = client.Mail(p.config.From); err != nil {
		p.log.Error(err)
		return err
	}

	if err = client.Rcpt(to); err != nil {
		p.log.Error(err)
		return err
	}

	// Data
	w, err := client.Data()
	if err != nil {
		p.log.Error(err)
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		p.log.Error(err)
		return err
	}

	err = w.Close()
	if err != nil {
		p.log.Error(err)
		return err
	}

	err = client.Quit()
	if err != nil {
		p.log.Error(err)
		return err
	}

	p.log.Info("Mail sent successfully")
	return nil
}

type SMTPMailProvider interface {
	Send(to string, subject string, message string) error
}

func NewSMTPMailProvider(log *zap.SugaredLogger, config *config.MailConfig) SMTPMailProvider {
	return &smtpMailProvider{log, config}
}
