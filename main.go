package main

import (
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	pass := flag.String("pass", "", "a string")
	service := flag.String("service", "", "a string")
	modifier := flag.String("modifier", "", "a string")

	flag.Parse()

	if *service == "" {
		fmt.Println("Service must be set.")
		return
	}

	if *pass == "" {
		fmt.Println("Your password: ")
		bytes, _ := terminal.ReadPassword(int(syscall.Stdin))
		*pass = string(bytes)
	}

	sha512Hash := generateHash(*service, *modifier, *pass)

	fmt.Printf("Generated password for %s@%s: %s\n", *modifier, *service, sha512Hash)
}

func generateHash(service, mod, pass string) string {
	h := sha512.New()
	h.Write([]byte(fmt.Sprintf("%s%s%s", service, mod, pass)))
	return hex.EncodeToString(h.Sum(nil))[:32]
}
