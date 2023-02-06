package token

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

/*
{
  "alg": "HS256",
  "typ": "JWT"
}
*/

type Payload struct {
	Name string `json:"name"`
	Role int    `json:"role"`
}

/*
{
  "name": "Mary Sue",
  "role": "1"
}
*/
