package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func SecretHash(username, clientId, clientSecret string) (string, error) {
	mac := hmac.New(sha256.New, []byte(clientSecret))

	_, err := mac.Write([]byte(username + clientId))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}
