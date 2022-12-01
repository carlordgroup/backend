package card

import (
	"carlord/data"
	"carlord/ent"
	"carlord/ent/account"
	"carlord/ent/card"
	"carlord/ent/user"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get Card godoc
// @Tags card
// @Summary get cards of a user
// @Produce json
// @Success 200 {object} []ent.Card
// @Router /card/ [get]
func (s *service) get(ctx *gin.Context) (int, any) {
	cards, err := s.client.Card.Query().Where(card.HasOwnerWith(user.HasAccountWith(account.ID(ctx.GetInt("id"))))).All(ctx)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, cards
}

// helper method
func (s *service) hasCard(ctx *gin.Context) (*ent.Card, error) {
	type cardID struct {
		ID int `uri:"id"`
	}
	var c cardID
	err := ctx.ShouldBindUri(&c)
	if err != nil {
		return nil, err
	}

	editCard, err := s.client.Card.Query().Where(card.And(card.HasOwnerWith(user.ID(ctx.MustGet("account").(*data.Account).ID)), card.ID(c.ID))).Only(ctx)
	return editCard, nil
}

type cardBinding struct {
	Number         string `json:"number" binding:"numeric,required"`
	CardholderName string `json:"cardholder_name" binding:"required"`
	ValidUntil     string `json:"valid_until" binding:"required"`
}

// Edit Card godoc
// @Tags card
// @Param request body cardBinding true "edit card"
// @Summary edit a card
// @Accept json
// @Produce json
// @Success 200 {object} ent.Card
// @Router /card/:id [post]
func (s *service) post(ctx *gin.Context) (int, any) {
	// check if the user does own this card
	oldCard, err := s.hasCard(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	if oldCard == nil {
		return http.StatusNotFound, errors.New("this is not your card")
	}
	var cardInfo cardBinding

	err = ctx.ShouldBindJSON(&cardInfo)
	if err != nil {
		return http.StatusBadRequest, err
	}
	oldCard, err = s.client.Card.UpdateOne(oldCard).SetNumber(cardInfo.Number).SetCardholderName(cardInfo.CardholderName).SetValidUntil(cardInfo.ValidUntil).Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, oldCard

}

// Create Card godoc
// @Tags card
// @Param request body cardBinding true "create card"
// @Summary create a card
// @Accept json
// @Produce json
// @Success 201 {object} ent.Card
// @Router /card/ [post]
func (s *service) create(ctx *gin.Context) (int, any) {

	var cardInfo cardBinding

	err := ctx.ShouldBindJSON(&cardInfo)
	if err != nil {
		return http.StatusBadRequest, err
	}
	newCard, err := s.client.Card.Create().SetNumber(cardInfo.Number).
		SetCardholderName(cardInfo.CardholderName).
		SetValidUntil(cardInfo.ValidUntil).
		SetOwnerID(s.client.User.Query().Where(user.HasAccountWith(account.ID(ctx.GetInt("id")))).OnlyIDX(ctx)).Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusCreated, newCard
}

// Delete Card godoc
// @Tags card
// @Summary delete a card
// @Accept json
// @Produce json
// @Success 204
// @Router /card/:id [delete]
func (s *service) delete(ctx *gin.Context) (int, any) {
	oldCard, err := s.hasCard(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	if oldCard == nil {
		return http.StatusNotFound, errors.New("this is not your card")
	}

	err = s.client.Card.DeleteOne(oldCard).Exec(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusNoContent, nil
}
