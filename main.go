/*
	PacktPublishing offers a free tech e-book each day on its website.
	The program below is able to display the last offer to the stdout
	or send an email to the specified recipient using the sender's credentials (gmail only).
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	packtURI            = "https://www.packtpub.com/packt/offers/free-learning"
	notificationSubject = "Pack Publishing Today's free ebook"
	to, from, pass      string
)

func main() {
	flag.StringVar(&to, "to", "", "use when you want to send the result message to the specified email: -to user@host.com")
	flag.StringVar(&from, "from", "", "send from the specified gmail acount: -from sender@host.com")
	flag.StringVar(&pass, "pass", "", "the sender's gmail password: -pass yousecretpass")
	flag.Parse()

	if len(to) > 0 {
		if len(from) == 0 && len(pass) == 0 {
			log.Fatal("to send email notification the sender's email address and password should be provided")
		}

		validTo, err := mail.ParseAddress(to)
		if err != nil {
			log.Fatal("badly formatted email address: ", err)
		}
		validFrom, err := mail.ParseAddress(from)
		if err != nil {
			log.Fatal("badly formatted email address: ", err)
		}

		msg := fetch()

		err = send(validTo.Address, validFrom.Address, pass, notificationSubject, msg)
		if err != nil {
			log.Fatal("sending error: ", err)
		}

		log.Printf("email sent from %v to %v\n", validFrom.Address, validTo.Address)
	} else {
		msg := fetch()
		fmt.Println(msg)
	}
}

// fetch requests the Pack Publishing web page with freely available tech e-book
// and extracts the title of the book returning it as a formatted string.
func fetch() string {
	doc, err := goquery.NewDocument(packtURI)
	if err != nil {
		log.Fatal(err)
	}

	t := doc.Find(".dotd-title h2").Text()
	r := fmt.Sprintf("%v. Check on http://bit.ly/1Bp3uB2", strings.TrimSpace(t))
	return r
}

// send prepares the email message and sends it to the specified recipient from the sender's gmail account
// with provided password.
func send(to, from, pass, subject, body string) error {
	host := "smtp.gmail.com"
	port := ":587"
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", from, pass, host)

	err := smtp.SendMail(host+port, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
