package main

import (
	"fmt"
	"os"
	"strings"
	"golang.org/x/term"
)

func PasswordPrompt() (string, error) {
	fmt.Print(`Password: `)
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println() // cursor does not automatically advance
	if err != nil {
		return "", fmt.Errorf("error reading password: %w", err)
	}

	stringPassword := strings.TrimSpace(string(bytePassword))
	return stringPassword, nil
}

func main() {
	passwd, err := PasswordPrompt()	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("submitted password: '%s', length = %d\n", passwd, len(passwd))
}
