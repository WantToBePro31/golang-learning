package main

import "fmt"

func EmailInfo(email string) string {
	var (
		domainMode bool
		tldMode bool
	)
	domain := ""
	tld := ""
	for _, val := range email {
		if domainMode && val != '.'{
			domain += string(val)
		}
		if tldMode {
			tld += string(val)
		}
		if val == '@' {
			domainMode = true
		}
		if val == '.' && !tldMode {
			domainMode = false
			tldMode = true
		}
	}
	return fmt.Sprintf("Domain: %s dan TLD: %s", domain, tld) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
