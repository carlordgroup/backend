package card

import (
	"carlord/ent"
	"carlord/ent/card"
	"carlord/ent/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type service struct {
	client *ent.Client
}

type Authenticate interface {
	MustLogin() gin.HandlerFunc
}

func New(client *ent.Client) *service {
	s := service{
		client: client,
	}
	return &s
}

func (s *service) RegisterRouter(group gin.IRoutes, auth Authenticate) {
	group.Use(auth.MustLogin())
	group.GET("/", s.get)
	group.POST("/:id", s.post)
	group.DELETE(auth.GetAccountUser())
}

func (s *service) get(ctx *gin.Context) {
	cards, err := s.client.Card.Query().Where(card.HasOwnerWith(user.ID(ctx.MustGet("id").(int)))).All(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, cards)

}

func (s *service) post(ctx *gin.Context) {
	type cardID struct {
		ID int `uri:"id"`
	}
	var c cardID
	err := ctx.ShouldBindUri(&c)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var cardInfo ent.Card
	err = ctx.ShouldBindJSON(&cardInfo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	editCard, err := s.client.Card.Query().Where(card.And(card.HasOwnerWith(user.ID(ctx.MustGet("id").(int))), card.ID(c.ID))).Only(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	editCard, err = s.client.Card.UpdateOne(editCard).SetNumber(cardInfo.Number).SetCardholderName(cardInfo.CardholderName).SetValidUntil(cardInfo.ValidUntil).Save(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, editCard)

}
