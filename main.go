package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"net/smtp"
)

func main() {
	var arg_smtp_from string
	var arg_smtp_hello string
	var arg_smtp_port string
	var arg_smtp_server string

	flag.StringVar(&arg_smtp_from, "f", "user@example.com", "SMTP user (used in MAIL FROM)")
	flag.StringVar(&arg_smtp_hello, "h", "example.com", "SMTP hello (used in HELO/EHLO)")
	flag.StringVar(&arg_smtp_port, "p", "smtp", "SMTP port")
	flag.StringVar(&arg_smtp_server, "s", "", "SMTP hostname")
	flag.Parse()

	arg_rcpts := flag.Args()
	if (len(arg_rcpts) <= 0) {
		fmt.Printf("Usage:\n  %s [arguments] recipients...\n\nArguments:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	for _, arg_rcpt := range arg_rcpts {
		var smtp_server string

		if arg_smtp_server == "" {
			rcpt_domain := strings.Split(arg_rcpt, "@")[1]
			mxs, err := net.LookupMX(rcpt_domain)
			if err == nil {
				smtp_server = mxs[0].Host
			} else {
				smtp_server = rcpt_domain
			}
		} else {
			smtp_server = arg_smtp_server
		}

		fmt.Printf("Dialing to [%s]\n", smtp_server)
		client, err := smtp.Dial(smtp_server + ":" + arg_smtp_port)

		fmt.Printf("HELO/EHLO [%s]\n", arg_smtp_hello)
		err = client.Hello(arg_smtp_hello)
		fmt.Printf("err = [%s]\n", err)

		fmt.Printf("MAIL FROM: [%s]\n", arg_smtp_from)
		err = client.Mail(arg_smtp_from)
		fmt.Printf("err = [%s]\n", err)

		fmt.Printf("RCPT TO: [%s]\n", arg_rcpt)
		err = client.Rcpt(arg_rcpt)
		fmt.Printf("err = [%s]\n", err)

		client.Close()
	}
}
