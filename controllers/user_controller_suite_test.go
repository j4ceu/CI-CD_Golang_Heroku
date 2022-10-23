package controllers

import (
	"CICD_GolangtoHeroku/dto"
	"CICD_GolangtoHeroku/models"
	"CICD_GolangtoHeroku/services/mock"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteUser struct {
	suite.Suite
	controller UserController
	mock       *mock.UserServiceMock
}

func (s *suiteUser) SetupSuite() {
	mock := new(mock.UserServiceMock)
	s.mock = mock
	s.controller = NewUserController(s.mock)
}

func (s *suiteUser) TestCreateUser() {

	userJSON := `{"email":"jace@gmail.com", "password":"123456"}`

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedBody       models.User
	}{
		{

			Name:               "Test Create User",
			ExpectedStatusCode: 200,
			Method:             "POST",
			HasReturnBody:      true,
			ExpectedBody: models.User{
				Email:    "jace@gmail.com",
				Password: "123456",
			},
		},
	}

	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.mock.On("CreateUser", tc.ExpectedBody).Return(models.User{Email: "jace@gmail.com", Password: "123456"}, nil)
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			err := s.controller.CreateUser(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.HasReturnBody {
				var bodyResponse map[string]interface{}
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Email, bodyResponse["data"].(map[string]interface{})["email"])
				s.Equal(tc.ExpectedBody.Password, bodyResponse["data"].(map[string]interface{})["password"])

			}
		})

	}
}

func (s *suiteUser) TestCreateUserError400() {

	userJSON := `{"email":"jace@gmail.com", "password":123456}`

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedBody       map[string]interface{}
	}{
		{

			Name:               "Test Create User Error 400",
			ExpectedStatusCode: 400,
			Method:             "POST",
			HasReturnBody:      true,
			ExpectedBody: map[string]interface{}{
				"error": "code=400, message=Unmarshal type error: expected=string, got=number, field=password, offset=44, internal=json: cannot unmarshal number into Go struct field User.password of type string",
			},
		},
	}

	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			err := s.controller.CreateUser(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.HasReturnBody {
				var bodyResponse map[string]interface{}
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody["error"], bodyResponse["error"])

			}
		})

	}
}

func (s *suiteUser) TestGetAllUsers() {
	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedBody       []models.User
	}{
		{

			Name:               "Test Get All Users",
			ExpectedStatusCode: 200,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedBody: []models.User{
				{
					Email:    "jace@gmail.com",
					Password: "123456",
				},
			},
		},
	}

	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.mock.On("GetAllUsers").Return([]models.User{
				{
					Email:    "jace@gmail.com",
					Password: "123456",
				},
			}, nil).Once()

			req := httptest.NewRequest(tc.Method, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			err := s.controller.GetAllUsers(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.HasReturnBody {
				var bodyResponse map[string]interface{}
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody[0].Email, bodyResponse["data"].([]interface{})[0].(map[string]interface{})["email"])
				s.Equal(tc.ExpectedBody[0].Password, bodyResponse["data"].([]interface{})[0].(map[string]interface{})["password"])
			}
		})
	}
}

func (s *suiteUser) TestGetAllUsersError500() {
	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedBody       map[string]interface{}
	}{
		{

			Name:               "Test Get All Users Error 500",
			ExpectedStatusCode: 500,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedBody: map[string]interface{}{
				"error": "Internal Server Error",
			},
		},
	}

	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.mock.On("GetAllUsers").Return([]models.User{}, errors.New("Internal Server Error")).Once()

			req := httptest.NewRequest(tc.Method, "/", nil)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			err := s.controller.GetAllUsers(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.HasReturnBody {
				var bodyResponse map[string]interface{}
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody["error"], bodyResponse["error"])
			}
		})
	}
}

func (s *suiteUser) TestLoginUser() {
	userJSON := `{"email":"jace@gmail.com", "password":"123456"}`
	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedBody       dto.UserResponse
	}{
		{

			Name:               "Test Login User",
			ExpectedStatusCode: 200,
			Method:             "POST",
			HasReturnBody:      true,
			ExpectedBody: dto.UserResponse{
				Email: "jace@gmail.com",
				Token: "123456",
			},
		},
	}

	for _, tc := range testCase {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.mock.On("LoginUser", "jace@gmail.com", "123456").Return(dto.UserResponse{Email: "jace@gmail.com", Token: "123456"}, nil)
			req := httptest.NewRequest(tc.Method, "/", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, res)

			err := s.controller.LoginUser(c)
			s.NoError(err)

			s.Equal(tc.ExpectedStatusCode, res.Code)

			if tc.HasReturnBody {
				var bodyResponse map[string]interface{}
				s.NoError(json.Unmarshal(res.Body.Bytes(), &bodyResponse))
				s.Equal(tc.ExpectedBody.Email, bodyResponse["data"].(map[string]interface{})["email"])
				s.Equal(tc.ExpectedBody.Token, bodyResponse["data"].(map[string]interface{})["token"])

			}
		})

	}
}

func TestSuiteUser(t *testing.T) {
	suite.Run(t, new(suiteUser))

}
