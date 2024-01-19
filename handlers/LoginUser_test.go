package handlers

import (
	"errors"
	"gobus/dto"

	"github.com/go-playground/validator/v10"

	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LoginUser(t *testing.T) {
	test := []struct {
		Name           string
		Body           *dto.LoginRequest
		Route          string
		ExpectedError  error
		ExpectedResult string
	}{
		{
			Name: "Test-1",
			Body: &dto.LoginRequest{
				Email:    "dasdasd",
				Password: "1234",
			},
			ExpectedError: nil,
			ExpectedResult: "1234dr",
		},
	}
	for _, tc := range test {
		t.Run(tc.Name, func(t *testing.T) {
			v := validator.New()
			v.Struct(tc.Body)

			tes := func(login *dto.LoginRequest) (map[string]string, error) {
				return nil, nil
			}
			tes(tc.Body)
			if tc.ExpectedError == nil {
				require.Equal(t, tc.ExpectedResult, "1234ddr")
			} else {
				require.Equal(t, tc.ExpectedError, errors.New("error here"))
			}
		})
	}
}
