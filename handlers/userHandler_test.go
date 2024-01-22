package handlers

import (
	"encoding/json"
	"gobus/dto"
	"gobus/entities"
	"gobus/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
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

func TestUserHandler_FindBus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	test := []struct {
		name        string
		body        *dto.BusRequest
		beforeTest  func(userService *services.MockUserService)
		route       string
		errorResult map[string]interface{}
	}{
		{
			name: "success",
			body: &dto.BusRequest{
				DepartureStation: "Kannur",
				ArrivalStation:   "Bangalore",
			},
			// {BusID: 1, BusNumber: "KL18AA5555", TotalSleeperSeats: 10, TotalPushBackSeats: 10, BusTypeCode: "ACSL"}
			beforeTest: func(userService *services.MockUserService) {
				userService.EXPECT().FindBus(&dto.BusRequest{
					DepartureStation: "Kannur",
					ArrivalStation:   "Bangalore",
				}).Return([]*entities.BusesResp{}, nil)
			},
			route:       "/user/findbus",
			errorResult: nil,
		},
		{
			name: "fail case",
			body: &dto.BusRequest{
				DepartureStation: "Kannur",
				ArrivalStation:   "",
			},
			// {BusID: 1, BusNumber: "KL18AA5555", TotalSleeperSeats: 10, TotalPushBackSeats: 10, BusTypeCode: "ACSL"}
			beforeTest: func(userService *services.MockUserService) {
				userService.EXPECT().FindBus(&dto.BusRequest{
					DepartureStation: "Kannur",
					ArrivalStation:   "",
				}).Return([]*entities.BusesResp{}, nil)
			},
			route:       "/user/findbus",
			errorResult: map[string]interface{}{"data": interface{}(nil), "message": "Stations cannot be empty.", "status": "Failed"},
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

			req, err := http.NewRequest(http.MethodGet, tc.route, r)
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

func TestUserHandler_AddPassenger(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	test := []struct {
		name        string
		body        *entities.PassengerInfo
		beforeTest  func(userService *services.MockUserService)
		route       string
		errorResult map[string]interface{}
	}{
		{
			name: "success",
			body: &entities.PassengerInfo{
				PassengerID: 1,
				Name:        "ABC",
				Age:         20,
				Gender:      "Male",
				UserID:      1,
			},
			beforeTest: func(userService *services.MockUserService) {
				userService.EXPECT().AddPassenger(&entities.PassengerInfo{
					PassengerID: 1,
					Name:        "ABC",
					Age:         20,
					Gender:      "Male",
					UserID:      1,
				}, "xyz@gmail.com").Return(&entities.PassengerInfo{PassengerID: 1,
					Name:   "ABC",
					Age:    20,
					Gender: "Male",
					UserID: 1}, nil)
			},
			route:       "/user/addpassenger",
			errorResult: nil,
		},
		{
			name: "fail case",
			body: &entities.PassengerInfo{
				PassengerID: 1,
				Name:        "",
				Age:         20,
				Gender:      "Male",
				UserID:      1,
			},
			beforeTest: func(userService *services.MockUserService) {
				userService.EXPECT().AddPassenger(&entities.PassengerInfo{
					PassengerID: 1,
					Name:        "",
					Age:         20,
					Gender:      "Male",
					UserID:      1,
				}, "xyz@gmail.com").Return(&entities.PassengerInfo{PassengerID: 1,
					Name:   "ABC",
					Age:    20,
					Gender: "Male",
					UserID: 1}, nil)
			},
			route:       "/user/addpassenger",
			errorResult: map[string]interface{}{"data": interface{}(nil), "message": "Missing mandatory fields.", "status": "Failed"},
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
				require.Equal(t, w.Code, 201)
			}
		})
	}
}

// func TestUserHandler_FindBus(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	test := []struct {
// 		name        string
// 		body        *dto.BusRequest
// 		beforeTest  func(userService *services.MockUserService)
// 		route       string
// 		errorResult map[string]interface{}
// 	}{
// 		{
// 			name: "success",
// 			body: &dto.BusRequest{
// 				DepartureStation: "Kannur",
// 				ArrivalStation:   "Bangalore",
// 			},
// 			// {BusID: 1, BusNumber: "KL18AA5555", TotalSleeperSeats: 10, TotalPushBackSeats: 10, BusTypeCode: "ACSL"}
// 			beforeTest: func(userService *services.MockUserService) {
// 				userService.EXPECT().FindBus(&dto.BusRequest{
// 					DepartureStation: "Kannur",
// 					ArrivalStation:   "Bangalore",
// 				}).Return([]*entities.BusesResp{}, nil)
// 			},
// 			route:       "/user/findbus",
// 			errorResult: nil,
// 		},
// 		{
// 			name: "fail case",
// 			body: &dto.BusRequest{
// 				DepartureStation: "Kannur",
// 				ArrivalStation:   "",
// 			},
// 			// {BusID: 1, BusNumber: "KL18AA5555", TotalSleeperSeats: 10, TotalPushBackSeats: 10, BusTypeCode: "ACSL"}
// 			beforeTest: func(userService *services.MockUserService) {
// 				userService.EXPECT().FindBus(&dto.BusRequest{
// 					DepartureStation: "Kannur",
// 					ArrivalStation:   "",
// 				}).Return([]*entities.BusesResp{}, nil)
// 			},
// 			route:       "/user/findbus",
// 			errorResult: map[string]interface{}{"data": interface{}(nil), "message": "Stations cannot be empty.", "status": "Failed"},
// 		},
// 	}
// 	for _, tc := range test {
// 		t.Run(tc.name, func(t *testing.T) {
// 			MockUserService := services.NewMockUserService(ctrl)
// 			u := &UserHandler{
// 				user: MockUserService,
// 			}
// 			if tc.beforeTest != nil {
// 				tc.beforeTest(MockUserService)
// 			}
// 			gin.SetMode(gin.TestMode)
// 			w := httptest.NewRecorder()
// 			engine := gin.Default()

// 			RegisterUserRoutes(engine, u)

// 			body, err := json.Marshal(tc.body)
// 			if err != nil {
// 				require.NoError(t, err)
// 			}

// 			r := strings.NewReader(string(body))

// 			req, err := http.NewRequest(http.MethodGet, tc.route, r)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			engine.ServeHTTP(w, req)
// 			if tc.errorResult != nil {
// 				errValue, _ := json.Marshal(tc.errorResult)
// 				require.JSONEq(t, w.Body.String(), string(errValue))
// 			} else {
// 				require.Equal(t, w.Code, 202)
// 			}
// 		})
// 	}
// }

func TestUserHandler_ViewAllPassengers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	test := []struct {
		name        string
		body        string
		beforeTest  func(userService *services.MockUserService)
		route       string
		errorResult map[string]interface{}
	}{
		{
			name: "success",
			body: "abc@gmail.com",
			beforeTest: func(userService *services.MockUserService) {
				userService.EXPECT().ViewAllPassengers("abc@gmail.com").Return([]*entities.PassengerInfo{}, nil)
			},
			route:       "/user/viewallpassenger",
			errorResult: nil,
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

			req, err := http.NewRequest(http.MethodGet, tc.route, r)
			if err != nil {
				t.Fatal(err)
			}
			engine.ServeHTTP(w, req)
			if tc.errorResult != nil {
				errValue, _ := json.Marshal(tc.errorResult)
				require.JSONEq(t, w.Body.String(), string(errValue))
			} else {
				require.Equal(t, w.Code, 302)
			}
		})
	}
}
func TestUserHandler_BookSeat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	test := []struct {
		name        string
		body        *dto.BookingRequest
		beforeTest  func(userService *services.MockUserService)
		route       string
		errorResult map[string]interface{}
	}{
		{
			name: "success",
			body: &dto.BookingRequest{BusID: 1, PassengerID: pq.Int64Array{1, 2}, SeatsReserved: []string{"01A", "01B"}, BookingDate: "01012024"},
			beforeTest: func(userService *services.MockUserService) {
				userService.EXPECT().BookSeat(&dto.BookingRequest{BusID: 1, PassengerID: pq.Int64Array{1, 2}, SeatsReserved: []string{"01A", "01B"}, BookingDate: "01012024"}, "abc@gmail.com").Return(&entities.Booking{}, nil)
			},
			route:       "/user/bookseat",
			errorResult: nil,
		},
		{
			name: "fail case",
			body: &dto.BookingRequest{BusID: 1, PassengerID: pq.Int64Array{1, 2}, SeatsReserved: []string{"01A", "01B"}, BookingDate: ""},
			beforeTest: func(userService *services.MockUserService) {
				userService.EXPECT().BookSeat(&dto.BookingRequest{BusID: 1, PassengerID: pq.Int64Array{1, 2}, SeatsReserved: []string{"01A", "01B"}, BookingDate: ""}, "abc@gmail.com").Return(&entities.Booking{}, nil)
			},
			route:       "/user/bookseat",
			errorResult: map[string]interface{}{"data": interface{}(nil), "message": "Mandatory fields cannot be empty", "status": "Failed"},
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
