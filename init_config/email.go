package init_config

import (
	"net/smtp"
	"strconv"
)

var Eamil *emailClient

func init() {
	auth := smtp.PlainAuth("", EmailConf().Account, EmailConf().Password, EmailConf().Host)
	client := &emailClient{}
	client.Auth = auth
	client.emailConf = EmailConf()
	Eamil = client
}

type emailClient struct {
	smtp.Auth
	emailConf
}

func (client *emailClient) SendMail(to []string, message string) error {
	err := smtp.SendMail(client.Host+":"+strconv.Itoa(client.Port), client.Auth, client.Account, to, []byte(message))
	return err
}
