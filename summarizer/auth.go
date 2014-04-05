package summarizer

import (
	"fmt"
	"github.com/gcmurphy/getpass"
	"github.com/k0kubun/gothub"
	"log"
	"os"
	"os/user"
)

func github() *gothub.GitHub {
	guestClient, _ := gothub.Guest()
	if guestClient != nil {
		return guestClient
	}

	if !authenticated() {
		fmt.Println("Guest API Limit exceeded.")
		authenticate()
	}

	loginClient, err := gothub.BasicLogin(usernameAndPassword())
	if err != nil {
		deauthenticate()

		//FIXME: Support Two Factor Authentication
		fmt.Println("Login failed. It seems that password is wrong or you use two factor authentication.")
		fmt.Println("I'm sorry but two factor authentication is not supported now.")
		log.Fatal(err)
	}

	return loginClient
}

func configFilePath() (path string) {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	path = fmt.Sprintf("%s/%s", user.HomeDir, ".rockstar")
	return
}

func usernameAndPassword() (username, password string) {
	file, err := os.OpenFile(configFilePath(), os.O_RDONLY, 0)
	if err != nil {
		return "", ""
	}
	defer func() {
		file.Close()
	}()

	fmt.Fscanf(file, "%s\n", &username)
	fmt.Fscanf(file, "%s\n", &password)
	return
}

func authenticated() (authenticated bool) {
	username, password := usernameAndPassword()
	authenticated = (username != "") && (password != "")
	return
}

func authenticate() {
	var username, password string
	fmt.Println("Enter your GitHub account credentials.")

	fmt.Printf("Username: ")
	fmt.Scanf("%s", &username)
	password, _ = getpass.GetPass()
	fmt.Println()

	file, err := os.OpenFile(configFilePath(), os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file.Close()
	}()

	fmt.Fprintf(file, "%s\n", username)
	fmt.Fprintf(file, "%s\n", password)
}

func deauthenticate() {
	os.Remove(configFilePath())
}
