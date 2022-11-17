package management

import (
	"carlord/ent"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func (s *service) deleteLocation(ctx *gin.Context, id int) (int, any) {

	err := s.client.Location.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusNoContent, nil
}

func (s *service) listLocation(ctx *gin.Context) (int, any) {

	locations, err := s.client.Location.Query().All(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, locations
}
