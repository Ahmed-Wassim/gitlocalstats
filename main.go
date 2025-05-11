package main

import "flag"

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "folder", "", "add a new folder to scan for Git Repositories")
	flag.StringVar(&email, "email", "ahmedwassim317@gmail.com", "email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
