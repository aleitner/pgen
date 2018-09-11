package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("pgen", map[string]interface{}{
		"generateHash": generateHash,
	})
}

func generateHash(service, mod, salt string, length int) string {
	h := sha512.New()
	h.Write([]byte(fmt.Sprintf("%s%s%s", salt, mod, service)))
	return hex.EncodeToString(h.Sum(nil))[:length]
}
