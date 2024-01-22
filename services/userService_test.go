package services

import (
	"errors"
	"gobus/dto"
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
	type args struct {
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
					nil,
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
			wantErr: false,
		},
		{
			name: "success register user",
			args: args{
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

func Test_ViewBookings(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(userRepo *repository.MockUserRepository)
		want       []*entities.Booking
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				input: "abc@gmail.com",
			},
			beforeTest: func(userRepo *repository.MockUserRepository) {
				userRepo.EXPECT().ViewBookings("abc@gmail.com").Return([]*entities.Booking{{UserID: 1}},
					nil,
				)
			},
			want:    []*entities.Booking{{UserID: 1}},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				input: "abc@gmail.com",
			},
			beforeTest: func(userRepo *repository.MockUserRepository) {
				userRepo.EXPECT().ViewBookings("abc@gmail.com").Return([]*entities.Booking{{UserID: 1}},
					errors.New("Oops"),
				)
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

			got, err := w.ViewBookings(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("services.ViewBookings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("services.ViewBookings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AddPassenger(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		passenger *entities.PassengerInfo
		email     string
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(userRepo *repository.MockUserRepository)
		want       *entities.PassengerInfo
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				passenger: &entities.PassengerInfo{Name: "ABC", Age: 10, Gender: "Male"},
				email:     "abc@gmail.com",
			},
			beforeTest: func(userRepo *repository.MockUserRepository) {
				userRepo.EXPECT().AddPassenger(&entities.PassengerInfo{Name: "ABC", Age: 10, Gender: "Male"}, "abc@gmail.com").Return(&entities.PassengerInfo{Name: "ABC", Age: 10, Gender: "Male"},
					nil,
				)
			},
			want:    &entities.PassengerInfo{Name: "ABC", Age: 10, Gender: "Male"},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				passenger: &entities.PassengerInfo{Name: "ABC", Age: 10, Gender: "Male"},
				email:     "abc@gmail.com",
			},
			beforeTest: func(userRepo *repository.MockUserRepository) {
				userRepo.EXPECT().AddPassenger(&entities.PassengerInfo{Name: "ABC", Age: 10, Gender: "Male"}, "abc@gmail.com").Return(&entities.PassengerInfo{Name: "ABC", Age: 10, Gender: "Male"},
					errors.New("Oops"),
				)
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

			got, err := w.AddPassenger(tt.args.passenger, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("services.AddPassenger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("services.AddPassenger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		LoginRequest *dto.LoginRequest
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(userRepo *repository.MockUserRepository)
		want       map[string]string
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				LoginRequest: &dto.LoginRequest{Email: "abc@gmail.com", Password: "1234"},
			},
			beforeTest: func(userRepo *repository.MockUserRepository) {
				userRepo.EXPECT().FindUserByEmail("abc@gmail.com").Return(&entities.User{Email: "abc@gmail.com", UserName: "abc", Password: "1234", Role: "user"},
					nil,
				)
			},
			want:    map[string]string{"access_token": "1234"},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				LoginRequest: &dto.LoginRequest{Email: "abc@gmail.com", Password: "1234"},
			},
			beforeTest: func(userRepo *repository.MockUserRepository) {
				userRepo.EXPECT().FindUserByEmail("abc@gmail.com").Return(&entities.User{Email: "abc@gmail.com", UserName: "abc", Password: "1234", Role: "user"},
					errors.New("Oops"),
				)
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

			got, err := w.Login(tt.args.LoginRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("services.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("services.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
