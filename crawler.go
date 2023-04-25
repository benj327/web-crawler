package main

import (
	//"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"net/mail"
	"regexp"
	"strings"
)

func findSensitiveInfo(url string, results chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	// Find email addresses
	doc.Find("body").Each(func(_ int, s *goquery.Selection) {
		textNodes := []string{}
		s.Contents().Each(func(_ int, n *goquery.Selection) {
			if n.Text() != "" {
				textNodes = append(textNodes, n.Text())
			}
		})

		text := strings.Join(textNodes, " ")
		emails := getEmails(text)
		for _, email := range emails {
			results <- email
		}
	})
}


func getEmails(text string) []string {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	matches := emailRegex.FindAllString(text, -1)
	emails := []string{}

	for _, match := range matches {
		addr, err := mail.ParseAddress(match)
		if err == nil {
			emails = append(emails, addr.Address)
		} else {
			// Handle concatenated email addresses
			emailParts := strings.Split(match, "@")
			if len(emailParts) == 2 {
				userPart := emailParts[0]
				domainPart := emailParts[1]

				// Extract the last part of the user portion
				userParts := strings.Split(userPart, " ")
				user := userParts[len(userParts)-1]

				// Extract the first part of the domain portion
				domainParts := strings.Split(domainPart, " ")
				domain := domainParts[0]

				email := user + "@" + domain
				if _, err := mail.ParseAddress(email); err == nil {
					emails = append(emails, email)
				}
			}
		}
	}

	return emails
}

