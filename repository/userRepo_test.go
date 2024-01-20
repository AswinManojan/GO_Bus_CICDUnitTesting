package repository

import (
	"errors"
	"gobus/entities"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_userRepo_Register(t *testing.T) {
	tests := []struct {
		Name       string
		args       *entities.User
		beforeTest func(sqlmock.Sqlmock)
		want       *entities.User
		wantErr    bool
	}{
		{
			Name: "Fail Register User",
			args: &entities.User{
				ID:          1,
				Email:       "aswinmanoj@gmail.com",
				UserName:    "Aswin Manoj",
				Password:    "1234",
				PhoneNumber: "1234567890",
				Gender:      "male",
				DOB:         "26121998",
			},
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				s.ExpectQuery(regexp.QuoteMeta(`
				INSERT INTO "users" ("email","user_name","password","role","phone_number","gender","dob","is_locked","user_wallet","id")
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
				RETURNING "id"`)).
					WithArgs("aswinmanoj@gmail.com", "Aswin Manoj", "1234", "user", "1234567890", "male", "26121998", false, 0, 1).
					WillReturnError(errors.New("Oops,error"))
				s.ExpectRollback()
			},
			wantErr: true,
		},
		// {
		// 	Name: "Success Register User",
		// 	args: &entities.User{
		// 		ID:          1,
		// 		Email:       "aswinmanoj@gmail.com",
		// 		UserName:    "Aswin Manoj",
		// 		Password:    "1234",
		// 		PhoneNumber: "1234567890",
		// 		Gender:      "male",
		// 		DOB:         "26121998",
		// 	},
		// 	beforeTest: func(s sqlmock.Sqlmock) {
		// 		s.ExpectBegin()
		// 		s.ExpectExec(regexp.QuoteMeta(`
		// 		INSERT INTO "users" ("email","user_name","password","role","phone_number","gender","dob","is_locked","user_wallet","id")
		// 		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		// 		RETURNING "id"`)).
		// 			WithArgs("aswinmanoj@gmail.com", "Aswin Manoj", "1234", "user", "1234567890", "male", "26121998", false, 0, 1).
		// 		// 	WillReturnError(errors.New("Oops,error"))
		// 		// s.ExpectRollback()
		// 		// s.ExpectExec(regexp.QuoteMeta(`INSERT INTO users (id,email,user_name,password,phone_number,gender,d_o_b)
		// 		// values ($1,$2,$3,$4,$5,$6,$7);`)).WithArgs(1, "aswinmanoj@gmail.com", "Aswin Manoj", "1234", "1234567890", "male", "26121998").
		// 		WillReturnResult(sqlmock.NewResult(1, 1))
		// 	},
		// 	want: &entities.User{
		// 		Email:       "aswinmanoj@gmail.com",
		// 		UserName:    "Aswin Manoj",
		// 		Password:    "1234",
		// 		PhoneNumber: "1234567890",
		// 		Gender:      "male",
		// 		DOB:         "26121998"},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			testdb, _ := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
			u := &UserRepositoryImpl{
				DB: testdb,
			}
			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}
			got, err := u.RegisterUser(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
