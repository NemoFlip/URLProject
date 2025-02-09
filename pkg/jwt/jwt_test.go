package jwt_test

import (
	"URLProject/pkg/jwt"
	"github.com/joho/godotenv"
	"os"

	"testing"
)

func TestJWT_Create(t *testing.T) {
	const email = "test@gmail.com"
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatal(err)
	}
	jwtToken := jwt.NewJWT(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := jwtToken.Create(jwt.JWTPayload{Email: email})
	if err != nil {
		t.Fatal(err)
	}

	isValid, jwtPayload := jwtToken.Parse(tokenString)
	if !isValid || jwtPayload.Email != email {
		t.Fatal("token is invalid")
	}

}
