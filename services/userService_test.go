package services

import (
	"context"
	"errors"
	"gobus/entities"
	"gobus/middleware"
	"gobus/repository"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_register_user(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// c := repository.NewMockUserRepository(ctrl)
	type args struct {
		ctx   context.Context
		input *entities.User
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(userRepo *repository.MockUserRepository)
		want       *entities.User
		wantErr    bool
	}{
		{
			name: "success register user",
			args: args{
				ctx: context.TODO(),
				input: &entities.User{
					ID:          1,
					Email:       "aswinmanoj@gmail.com",
					UserName:    "Aswin Manoj",
					Password:    "1234",
					PhoneNumber: "1234567890",
					Gender:      "male",
					DOB:         "26121998",
				},
			},
			beforeTest: func(userRepo *repository.MockUserRepository) {
				userRepo.EXPECT().RegisterUser(&entities.User{
					ID:          1,
					Email:       "aswinmanoj@gmail.com",
					UserName:    "Aswin Manoj",
					Password:    "1234",
					PhoneNumber: "1234567890",
					Gender:      "male",
					DOB:         "26121998",
				},
				).Return(&entities.User{
					ID:          1,
					Email:       "aswinmanoj@gmail.com",
					UserName:    "Aswin Manoj",
					Password:    "1234",
					PhoneNumber: "1234567890",
					Gender:      "male",
					DOB:         "26121998",
				},
					errors.New("oops"),
				)
			},
			want: &entities.User{
				ID:          1,
				Email:       "aswinmanoj@gmail.com",
				UserName:    "Aswin Manoj",
				Password:    "1234",
				PhoneNumber: "1234567890",
				Gender:      "male",
				DOB:         "26121998",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := repository.NewMockUserRepository(ctrl)

			w := &UserServiceImpl{
				repo: mockUserRepo,
				jwt:  middleware.NewJwtUtil(),
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockUserRepo)
			}

			got, err := w.RegisterUser(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("services.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("services.RegisterUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
