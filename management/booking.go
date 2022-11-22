package management

import (
	"github.com/gin-gonic/gin"
)

// List Booking godoc
// @Tags management
// @Summary list all bookings
// @Accept json
// @Produce json
// @Success 200 {object} []ent.Booking
// @Router /management/booking [get]
func (s *service) listBooking(ctx *gin.Context) (int, any) {
	data := s.client.Booking.Query().AllX(ctx)
	return 0, data
}
