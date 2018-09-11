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
	salt := flag.String("salt", "", "salt for generating pass")
	service := flag.String("service", "", "aservice for which pass is being created")
	modifier := flag.String("modifier", "", "Specific modifier about pass or service")
	length := flag.Int("length", 32, "Pass length")

	flag.Parse()

	if *service == "" {
		fmt.Println("Service must be set.")
		return
	}

	if *salt == "" {
		fmt.Println("Your password: ")
		bytes, _ := terminal.ReadPassword(int(syscall.Stdin))
		*salt = string(bytes)
	}

	sha512Hash := generateHash(*service, *modifier, *salt, *length)

	fmt.Printf("Generated password for %s@%s: %s\n", *modifier, *service, sha512Hash)
}

func generateHash(service, mod, salt string, length int) string {
	h := sha512.New()
	h.Write([]byte(fmt.Sprintf("%s%s%s", service, mod, salt)))
	return hex.EncodeToString(h.Sum(nil))[:length]
}
