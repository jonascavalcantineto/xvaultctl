package security

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func GenerateCredentialsPassword() string {
	fmt.Println("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err)
	}

	password := string(bytePassword)

	return strings.TrimSpace(password)
}
