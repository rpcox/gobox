package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"golang.org/x/term"
)

const (
// LdapURL       = `ldap://tequila.lan:389`
// BaseDN        = `ou=people,dc=cox,dc=lan`
// ServiceBindDN = `uid=minion,ou=service,dc=cox,dc=lan`
// ServiceBindPW = `Tomcat1200!`
)

func PasswordPrompt(prefix string) (string, error) {
	fmt.Print(prefix + ` password: `)
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println() // cursor does not automatically advance
	if err != nil {
		return "", fmt.Errorf("error reading password: %w", err)
	}

	stringPassword := strings.TrimSpace(string(bytePassword))
	return stringPassword, nil
}

func UserExists(li LdapInfo, saPasswd string) (bool, string, error) {

	serviceConn, err := ldap.DialURL(li.ldapUrl) // dial up the ldap server
	if err != nil {
		return false, "", fmt.Errorf("%%err: service account ldap connect: %w", err)
	}
	defer serviceConn.Close()

	err = serviceConn.Bind(li.serviceAccnt, saPasswd) // bind with service account creds
	if err != nil {
		return false, "", fmt.Errorf("%%err: bind: %w for user '%s'", err, li.serviceAccnt)
	}

	searchFilter := fmt.Sprintf("(uid=%s)", ldap.EscapeFilter(li.userAccnt))
	// https://pkg.go.dev/github.com/go-ldap/ldap/v3#SearchRequest
	searchRequest := ldap.NewSearchRequest(
		li.baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		searchFilter,
		[]string{"dn"},
		nil,
	)

	searchResult, err := serviceConn.Search(searchRequest)
	if err != nil {
		return false, "", fmt.Errorf("query for user=%s failed: %w", li.userAccnt, err)
	}

	if len(searchResult.Entries) == 0 {
		return false, "", nil
	} else if len(searchResult.Entries) > 1 {
		return false, "", fmt.Errorf("ambiguous return: %d users found", searchResult.Entries)
	}

	return true, searchResult.Entries[0].DN, nil

}

func UserAuthenticate(li LdapInfo, userPasswd string) error {
	userConn, err := ldap.DialURL(li.ldapUrl)
	if err != nil {
		return fmt.Errorf("%%err: user ldap connect: %w", err)
	}
	defer userConn.Close()

	err = userConn.Bind(li.userDN, userPasswd)
	if err != nil {
		if ldap.IsErrorWithCode(err, ldap.LDAPResultInvalidCredentials) {
			return fmt.Errorf("auth fail: %w for user '%s'", err, li.userAccnt)
		}
		return fmt.Errorf("%%err: bind: %w", err)
	}

	return nil
}

type LdapInfo struct {
	baseDN       string
	ldapUrl      string
	serviceAccnt string
	userAccnt    string
	userDN       string
}

func main() {
	var li LdapInfo
	flag.StringVar(&li.baseDN, "b", "ou=people,dc=cox,dc=lan", "Specify the base DN")
	flag.StringVar(&li.ldapUrl, "l", "ldap://tequila.lan:389", "LDAP destination ULR")
	flag.StringVar(&li.serviceAccnt, "s", "uid=minion,ou=service,dc=cox,dc=lan", "Specify the base DN")
	flag.StringVar(&li.userAccnt, "u", "", "Specify the uid of user to check")
	flag.Parse()

	//	pass := "ricldap"

	saPasswd, err := PasswordPrompt("service account")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	exists, dn, err := UserExists(li, saPasswd)
	if !exists && err == nil {
		fmt.Fprintf(os.Stdout, "user does not exist in directory information tree\n")
	} else if !exists && err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else if exists {
		li.userDN = dn
		fmt.Fprintf(os.Stderr, "user exists @ %s\n", dn)
		userPasswd, err := PasswordPrompt(li.userAccnt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%%err: %v", err)
			os.Exit(1)
		}

		err = UserAuthenticate(li, userPasswd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		fmt.Fprintln(os.Stdout, "user auth sucess")
	}

}
