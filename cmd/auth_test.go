package main

import (
	"URLProject/internal/delivery/payload"
	"URLProject/internal/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func initDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	db.Create(&entity.User{
		Name:     "test",
		Email:    "test@gmail.com",
		Password: string(hashedPassword),
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "test@gmail.com").
		Delete(&entity.User{})
}
func TestLoginSuccess(t *testing.T) {
	// Preparation
	db := initDB()
	initData(db)
	ts := httptest.NewServer(App())
	defer ts.Close()
	defer removeData(db)

	data, _ := json.Marshal(&payload.LoginRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	url := fmt.Sprintf("%s/auth/login", ts.URL)
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %d - got %d", http.StatusOK, resp.StatusCode)
	}
	var respStruct payload.LoginResponse
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response: %s", err)
	}
	if err = json.Unmarshal(respBody, &respStruct); err != nil {
		t.Fatalf("invalid response: %s", err)
	}
	if respStruct.Token == "" {
		t.Fatalf("token is empty")
	}

}

func TestLoginFail(t *testing.T) {
	db := initDB()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()
	defer removeData(db)

	data, _ := json.Marshal(&payload.LoginRequest{
		Email:    "gmail.com",
		Password: "string",
	})
	url := fmt.Sprintf("%s/auth/login", ts.URL)
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d - got %d", http.StatusBadRequest, resp.StatusCode)
	}
}
