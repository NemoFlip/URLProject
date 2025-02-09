package services

import (
	"URLProject/internal/entity"
	"testing"
)

type MockUserRepository struct {
}

func (repo *MockUserRepository) Create(user *entity.User) error {
	return nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*entity.User, error) {
	return nil, nil
}

func TestAuthService_Register(t *testing.T) {
	const email = "test@gmail.com"
	authService := NewAuthService(&MockUserRepository{})

	returnedEmail, err := authService.Register(email, "test", "Andrew")
	if err != nil {
		t.Fatal(err)
	}
	if returnedEmail != email {
		t.Fatalf("expected %s - got %s", email, returnedEmail)
	}
}
