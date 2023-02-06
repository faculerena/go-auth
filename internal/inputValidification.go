package data

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

func CheckUser() (b bool, err error) {

	username, password, err := getCredentials()
	if err != nil {
		return false, nil
	}

	onlyentry := data.GetUser(db, username)
	if onlyentry == nil {
		return false, nil
	}

	return onlyentry[0].ValidateHash(password), err
}

func getCredentials() (username string, password []byte, err error) {

	fmt.Print("> ")
	usernameRaw, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", nil, err
	}
	username = strings.TrimSuffix(usernameRaw, "\n")

	fmt.Print("> ")
	password, err = terminal.ReadPassword(0)
	if err != nil {
		return "", nil, err

	}

	return username, password, nil
}
