package handlers

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(c *gin.Engine, u *UserHandler) {
	c.POST("/user/login", u.Login)
	c.GET("/user/findbus", u.FindBus)
	c.POST("/user/addpassenger", u.AddPassenger)
	c.GET("/user/viewallpassenger", u.ViewAllPassengers)
	c.POST("/user/bookseat", u.BookSeat)
	c.GET("/user/coupon/view", u.FindCoupon)
	c.GET("/user/bookings/view", u.ViewBookings)
}
