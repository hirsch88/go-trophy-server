package config

func NewMailConfig() *MailConfig {
	return &MailConfig{
		/*
		|--------------------------------------------------------------------------
		| Mail Configurations
		|--------------------------------------------------------------------------
		|
		| In our case we decided to use the GMail SMTP Server. Just add your credentials
		| below or in the .env file.
		|
		*/

		Host:     Env("MAIL_HOST", "mail.cyon.ch"),
		Port:     Env("MAIL_PORT", "465"),
		Name:     Env("MAIL_NAME", "Go Trophy"),
		From:     Env("MAIL_FROM", "go-trophy@w3tec.ch"),
		Username: Env("MAIL_USERNAME", "go-trophy@w3tec.ch"),
		Password: Env("MAIL_PASSWORD", "password"),

	}
}

type MailConfig struct {
	Host     string
	Port     string
	Name     string
	From     string
	Username string
	Password string
}
