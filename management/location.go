package management

import (
	"carlord/ent"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Add Location godoc
// @Tags location
// @Param request body ent.Location true "location info"
// @Summary add a location
// @Accept json
// @Produce json
// @Success 201 {object} ent.Location
// @Router /management/location [post]
func (s *service) addLocation(ctx *gin.Context) (int, any) {
	var data ent.Location
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		return http.StatusBadRequest, err
	}

	newData, err := s.client.Location.Create().SetName(data.Name).SetLatitude(data.Latitude).SetLongitude(data.Longitude).Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusCreated, newData
}

// Update Location godoc
// @Tags location
// @Param request body ent.Location true "location info"
// @Summary update a location
// @Accept json
// @Produce json
// @Success 200 {object} ent.Location
// @Router /management/location/:id [post]
func (s *service) updateLocation(ctx *gin.Context, id int) (int, any) {
	var data ent.Location
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		return http.StatusBadRequest, err
	}

	newData, err := s.client.Location.UpdateOneID(id).SetName(data.Name).SetLatitude(data.Latitude).SetLongitude(data.Longitude).Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, newData
}

// Delete Location godoc
// @Tags location
// @Summary delete a location
// @Accept json
// @Success 204
// @Router /management/location [delete]
func (s *service) deleteLocation(ctx *gin.Context, id int) (int, any) {

	err := s.client.Location.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusNoContent, nil
}

// List Location godoc
// @Tags location
// @Summary add a location
// @Accept json
// @Produce json
// @Success 200 {object} []ent.Location
// @Router /management/location [get]
func (s *service) listLocation(ctx *gin.Context) (int, any) {

	locations, err := s.client.Location.Query().All(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, locations
}
