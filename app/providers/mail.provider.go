package providers

import (
	"github.com/hirsch88/go-trophy-server/app/mail"
	"github.com/hirsch88/go-trophy-server/config"
	"go.uber.org/zap"
)

type mailProvider struct {
	config   *config.MailConfig
	log      *zap.SugaredLogger
	template TemplateProvider
	smtpMail SMTPMailProvider
}

func (p *mailProvider) Send(mail mail.Mailable, to string) bool {
	p.log.Info("STARTING Send()")
	mailTemplate := mail.Build()
	message, err := p.template.Parse(mailTemplate.TemplatePath, mailTemplate.Context)
	if err != nil {
		p.log.Error("Could not parse mail template")
	}

	if err := p.smtpMail.Send(to, mailTemplate.Subject,  message); err != nil {
		p.log.Error("Could not send mail")
		return false
	}

	p.log.Info("FINISHED Send()")
	return true
}

type MailProvider interface {
	Send(mail mail.Mailable, to string) bool
}

func NewMailProvider(config *config.MailConfig, log *zap.SugaredLogger, template TemplateProvider, smtpMail SMTPMailProvider) MailProvider {
	return &mailProvider{config, log, template, smtpMail}
}
