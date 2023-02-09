# go-auth
Simple go auth service with JWT 

# How it works?

You can log in
```
curl http://localhost:9090/auth/signin -H 'Username:____' -H 'Passwordhash:_____'
```

You can sign in 
```
curl http://localhost:9090/auth/signup -H 'Username:____' -H 'Passwordhash:_____'
```

You can check if a token is valid! (By default, the token expires 1 minute after being created)
```
curl http://localhost:9090/auth/check -H 'Token:_____'
```





If you send a request with 'Username' and 'Passwordhash' as the header, the service will return the token if the credentials are valid.

### Token validation

Using this function you can listen when a request with header 'Token' reaches the service and check if the token is valid

```go
authRouter.Use(token.TokenValidationMiddleware)
```


# Setup

In the main folder, a 'private' folder exists these functions. Please adapt them to your necessity. These are needed to not hardcode
sensible data.

[WIP] change these to github secrets.

config.go
```go
package private
import "github.com/go-sql-driver/mysql"

func GetConfig() (config mysql.Config) {
    config = mysql.Config{
    User:   "user",
    Passwd: "password",
    Net:    "tcp",
    Addr:   "localhost:3306",
    DBName: "databasename",
    }
    return config
}
```

tokenHeader.go

```go
package private
func Header() (header string) {
	return "JWT SHA256"
}
```

tokenPayload.go

```go
package private

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

```

tokenSecret.go 

```go
package private

func Secret() string {
	
	return "your secret"
}
```

