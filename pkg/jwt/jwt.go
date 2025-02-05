package jwt

import "github.com/golang-jwt/jwt/v5"

type JWTPayload struct {
	Email string
}
type JWT struct {
	secretKey string
}

func NewJWT(secret string) *JWT {
	return &JWT{secretKey: secret}
}

func (j *JWT) Create(data JWTPayload) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})
	signedToken, err := jwtToken.SignedString([]byte(j.secretKey))
	if err != nil {
		 return "", err
	}
	return signedToken, nil
}

func (j *JWT) Parse(token string) (bool, *JWTPayload) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return false, nil
	}
	email, ok := jwtToken.Claims.(jwt.MapClaims)["email"].(string)
	if !ok {
		return false, nil
	}
	return jwtToken.Valid, &JWTPayload{Email: email}
}