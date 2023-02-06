package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/faculerena/goauth/private"
)

func GenerateToken(header string, payload map[string]string, secret string) (string, error) {

	secret = private.Secret()

	h := hmac.New(sha256.New, []byte(secret))
	header64 := base64.StdEncoding.EncodeToString([]byte(header))

	payloadstr, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error generating Token")
		return string(payloadstr), err
	}
	payload64 := base64.StdEncoding.EncodeToString(payloadstr)

	message := header64 + "." + payload64

	unsignedStr := header + string(payloadstr)

	h.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	tokenStr := message + "." + signature
	return tokenStr, nil
}
