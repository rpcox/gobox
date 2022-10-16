// simple mail client to test mail routing.  No passwd, no TLS
package main

import (
	"flag"
	"fmt"
	"log"

	mail "github.com/xhit/go-simple-mail/v2"
)

var htmlBody = `<!DOCTYPE html>
<html>
  <head><meta-equiv="Content-Type" content="text/html; charset=utf-8"></head>
    <body style='font-family:"Arial",sans-serif'>
      Mail routing success from %s
    </body>
</html>
`

var (
	Cc       = flag.String("cc", "", "Identify an account to CC (comma separated)")
	From     = flag.String("from", "no-reply@spud.com", "Identify the originating account")
	Subject  = flag.String("subject", "mail-test", "Identify the subject line")
	To       = flag.String("to", "rpcox@cox.net", "Identify the destination account")

	Port     = flag.Int("port", 25, "Identify the port used by mail server")
	Server   = flag.String("server", "smtp.cox.net", "Identify the mailhost")
)

func main() {

	server := mail.NewSMTPClient()
	server.Host = *Server
	server.Port = *Port 
	server.Authentication = mail.AuthNone

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	email := mail.NewMSG()
	email.SetFrom(*From)
	email.AddTo(*To)
	//email.AddCc("person@email.com")
	email.SetSubject(*Subject)

	email.SetBody(mail.TextHTML, fmt.Sprintf(htmlBody, *server))

	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal()
	}
}
