package privateExample

import "time"

func Payload(username string, passphrase string) (payload map[string]string) {
	currentTime := time.Now()
	expirationTime := currentTime.Add(time.Minute)

	payload = map[string]string{
		"usr": username,
		"psp": passphrase,
		"exp": expirationTime.Format(time.RFC3339),
	}

	return payload
}

/*

	payload = map[string]string{
		"Name": "Mary Sue",
		"Role": "0",
		"exp":  fmt.Sprint(time.Now().Add(time.Minute * 1).Unix())}

*/
