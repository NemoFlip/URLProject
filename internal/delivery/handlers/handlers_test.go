package handlers_test

import (
	"URLProject/configs"
	"URLProject/internal/delivery/handlers"
	"URLProject/internal/delivery/payload"
	"URLProject/internal/delivery/services"
	"URLProject/internal/repository"
	"URLProject/pkg/db"
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func bootstrap() (*handlers.AuthServer, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	dbGorm, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, mock, err
	}
	userRepository := repository.NewUserRepository(&db.Db{DB: dbGorm})
	cfg := configs.Config{
		Auth: configs.AuthConfig{SecretKey: "secret"},
	}
	authServer := handlers.NewAuthServer(handlers.AuthServerDeps{
		Config:      &cfg,
		AuthService: services.NewAuthService(userRepository),
	})
	return authServer, mock, nil
}

func TestLoginSuccess(t *testing.T) {
	authServer, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("test@gmail.com", string(hashedPassword))
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	data, _ := json.Marshal(&payload.LoginRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	authServer.LoginUser(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected %d - got %d", http.StatusOK, w.Code)
	}
}

func TestRegisterSuccess(t *testing.T) {
	authServer, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
	}

	rows := mock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	data, _ := json.Marshal(&payload.RegisterRequest{
		Name:     "Alex",
		Email:    "test@gmail.com",
		Password: "test",
	})
	reader := bytes.NewReader(data)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	authServer.RegisterUser(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected %d - got %d", http.StatusCreated, w.Code)
	}
}
