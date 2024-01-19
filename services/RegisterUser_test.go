package services

import (
	"errors"
	"gobus/entities"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RegisterUser(t *testing.T) {
	test := []struct {
		Name           string
		Body           *entities.User
		ExpectedResult *entities.User
		ExpectedError  error
	}{
		{
			Name: "Test-1",
			Body: &entities.User{
				Email:       "aswinn@gmail.com",
				UserName:    "aswinmanoj",
				Password:    "1234",
				PhoneNumber: "1234567890",
				Gender:      "male",
				DOB:         "26121998",
			},
			ExpectedResult: &entities.User{
				Email:       "aswinn@gmail.com",
				UserName:    "aswinmanoj",
				Password:    "1234",
				PhoneNumber: "1234567890",
				Gender:      "male",
				DOB:         "26121998",
			},
			ExpectedError: nil,
		},
		{
			Name: "Test-2",
			Body: &entities.User{
				Email:       "aswin@gmail.com",
				UserName:    "aswinmanoj",
				Password:    "1234",
				PhoneNumber: "1234567890",
				Gender:      "male",
				DOB:         "26121998",
			},
			ExpectedError: errors.New("user already exist"),
		},
	}
	for _, tc := range test {
		t.Run(tc.Name, func(T *testing.T) {
			Hspass = func(password string) (string, error) {
				return "", nil
			}
			tes_func := func(user *entities.User) (*entities.User, error) {
				test_entity := []*entities.User{
					{
						Email:       "aswin@gmail.com",
						UserName:    "aswinmanoj",
						Password:    "1234",
						PhoneNumber: "1234567890",
						Gender:      "male",
						DOB:         "26121998",
					},
				}
				for _, test := range test_entity {
					if test.Email == user.Email {
						return nil, errors.New("user already exist")
					}
				}
				return tc.Body, nil
			}
			tes_func(tc.Body)
			
			if tc.ExpectedError == nil {
				require.Equal(t, tc.ExpectedResult, tc.Body)
			} else {
				require.Equal(t, tc.ExpectedError, errors.New("user already exist"))
			}
		})
	}
}
