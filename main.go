package main

import (
	"flag"
	"fmt"
	"os"

	"net/smtp"
)

func main() {
	var smtp_from string
	var smtp_hello string
	var smtp_port string
	var smtp_server string

	flag.StringVar(&smtp_from, "f", "user@example.com", "SMTP user (used in MAIL FROM)")
	flag.StringVar(&smtp_hello, "h", "example.com", "SMTP hello (used in HELO/EHLO)")
	flag.StringVar(&smtp_port, "p", "smtp", "SMTP port")
	flag.StringVar(&smtp_server, "s", "smtp.google.com", "SMTP hostname")
	flag.Parse()

	rcpts := flag.Args()
	if (len(rcpts) <= 0) {
		fmt.Printf("Usage:\n  %s [arguments] recipients...\n\nArguments:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Dialing to " + smtp_server)
	client, err := smtp.Dial(smtp_server + ":" + smtp_port)

	fmt.Println("Hello: " + smtp_hello)
	err = client.Hello(smtp_hello)
	fmt.Println("err = ", err)

	fmt.Println("MAIL FROM: " + smtp_from)
	err = client.Mail(smtp_from)
	fmt.Println("err = ", err)

	for _, rcpt := range rcpts {
		fmt.Println("RCPT TO - " + rcpt)
		err = client.Rcpt(rcpt)
		fmt.Println("err = ", err)
	}

	client.Close()
}
