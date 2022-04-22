package controller

import (
	"app-echo/domain/model"
	"app-echo/service"

	"encoding/json"
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUser(t *testing.T) {
	e := echo.New()
	newUser := model.User{
		Email:    "alta@gmail.com",
		Password: "123456",
	}

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	us := service.NewMockUserService()
	us.Create(newUser)
	uc := NewUserController(us)
	uc.GetAll(c)

	var users []model.User
	if err := json.Unmarshal(rec.Body.Bytes(), &users); err != nil {
		t.Errorf("unmarshalling returned user failed")
	}

	assert.Len(t, users, 1)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, "alta@gmail.com", users[0].Email)
	assert.Equal(t, "123456", users[0].Password)
}

func TestCreateUser(t *testing.T) {
	e := echo.New()
	newUser := `{"email": "alta@gmail.com", "password": "123456"}`

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(newUser))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	us := service.NewMockUserService()
	uc := NewUserController(us)
	uc.Create(c)
	users, _ := us.GetAll()

	assert.Len(t, users, 1)
	assert.Equal(t, 201, rec.Code)
	assert.Equal(t, "alta@gmail.com", users[0].Email)
	assert.Equal(t, "123456", users[0].Password)
}
