package handlers

import (
	"encoding/json"
	"gobus/dto"
	"gobus/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	test := []struct {
		name        string
		body        *dto.LoginRequest
		beforeTest  func(userService *services.MockUserService)
		route       string
		errorResult map[string]interface{}
	}{
		{
			name: "success",
			body: &dto.LoginRequest{
				Email:    "aswin@gmail.com",
				Password: "1234",
			},
			beforeTest: func(userService *services.MockUserService) {
				userService.EXPECT().Login(&dto.LoginRequest{
					Email:    "aswin@gmail.com",
					Password: "1234",
				}).Return(map[string]string{"access_token": ""}, nil)
			},
			route:       "/user/login",
			errorResult: nil,
		},
		{
			name: "fail",
			body: &dto.LoginRequest{
				Email:    "aswin@gmail.com",
				Password: "",
			},
			beforeTest: func(userService *services.MockUserService) {
				userService.EXPECT().Login(&dto.LoginRequest{
					Email:    "aswin@gmail.com",
					Password: "",
				}).Return(map[string]string{"access_token": ""}, nil)
			},
			route:       "/user/login",
			errorResult: map[string]interface{}{"data": interface{}(nil), "message": "Password cannot be empty.", "status": "Failed"},
		},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			MockUserService := services.NewMockUserService(ctrl)
			u := &UserHandler{
				user: MockUserService,
			}
			if tc.beforeTest != nil {
				tc.beforeTest(MockUserService)
			}
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			engine := gin.Default()

			RegisterUserRoutes(engine, u)

			body, err := json.Marshal(tc.body)
			if err != nil {
				require.NoError(t, err)
			}

			r := strings.NewReader(string(body))

			req, err := http.NewRequest(http.MethodPost, tc.route, r)
			if err != nil {
				t.Fatal(err)
			}
			engine.ServeHTTP(w, req)
			if tc.errorResult != nil {
				errValue, _ := json.Marshal(tc.errorResult)
				require.JSONEq(t, w.Body.String(), string(errValue))
			} else {
				require.Equal(t, w.Code, 202)
			}
		})
	}
}
