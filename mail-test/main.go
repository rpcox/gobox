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
	From     = flag.String("from", "", "Identify the originating account")
	Subject  = flag.String("subject", "mail-test", "Identify the subject line")
	To       = flag.String("to", "", "Identify the destination account")

	Port     = flag.Int("port", 25, "Identify the port used by mail server")
	Server   = flag.String("server", "", "Identify the mailhost")
)

func main() {

	server := mail.NewSMTPClient()
	if len(*Server) > 0 {
		log.Fatal("-server must be identified")
	}
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

	if len(*Cc) > 0 {
		email.AddCc(*Cc)
	}
	email.SetSubject(*Subject)
	email.SetBody(mail.TextHTML, fmt.Sprintf(htmlBody, *server))

	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal()
	}
}
